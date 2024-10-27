import psycopg2
from psycopg2.extras import RealDictCursor, execute_batch
import numpy as np
from sentence_transformers import SentenceTransformer
from tqdm import tqdm
import uuid
import sys
import multiprocessing
from multiprocessing import Pool

psycopg2.extras.register_uuid()

def fetch_books(batch_size, offset):
    conn = psycopg2.connect(
        dbname='book_lens',
        user='postgres',
        password='password',
        host='localhost',
        port='5433'
    )
    cursor = conn.cursor(cursor_factory=RealDictCursor)
    query = """
    SELECT id, title, book_desc FROM books
    ORDER BY id
    LIMIT %s OFFSET %s;
    """
    cursor.execute(query, (batch_size, offset))
    books = cursor.fetchall()
    cursor.close()
    conn.close()
    return books

def store_embeddings(ids, embeddings):
    conn = psycopg2.connect(
        dbname='book_lens',
        user='postgres',
        password='password',
        host='localhost',
        port='5433'
    )
    cursor = conn.cursor()

    # Prepare data for insertion
    records = [(str(uuid.uuid4()), book_id, embedding.tolist()) for book_id, embedding in zip(ids, embeddings)]

    # Delete existing embeddings for these book_ids
    book_ids = [book_id for book_id in ids]
    delete_query = """
    DELETE FROM book_embeddings WHERE book_id = ANY(%s::uuid[]);
    """
    cursor.execute(delete_query, (book_ids,))

    # Use batch execution to insert new embeddings
    insert_query = """
    INSERT INTO book_embeddings (id, book_id, embedding)
    VALUES (%s, %s, %s);
    """

    execute_batch(cursor, insert_query, records)
    conn.commit()
    cursor.close()
    conn.close()

def generate_and_store_embeddings(batch_size=1000):
    model = SentenceTransformer('all-MiniLM-L6-v2')
    offset = 0
    total_processed = 0
    while True:
        books = fetch_books(batch_size, offset)
        if not books:
            break

        ids = []
        texts = []
        for book in books:
            book_id = book['id']
            title = book['title'] or ''
            description = book['book_desc'] or ''
            text = f"{title}. {description}"
            ids.append(book_id)
            texts.append(text)

        # Generate embeddings
        embeddings = model.encode(texts, batch_size=64, show_progress_bar=True)
        embeddings = embeddings / np.linalg.norm(embeddings, axis=1, keepdims=True)

        # Store embeddings
        store_embeddings(ids, embeddings)

        offset += batch_size
        total_processed += len(books)
        print(f"Processed {total_processed} books")



def generate_and_store_embeddings_parallel(num_workers=4, chunk_size=1000):
    book_ids = get_all_book_ids()
    total_books = len(book_ids)
    print(f"Total books to process: {total_books}")

    # Split the book IDs into chunks
    chunks = [book_ids[i:i + chunk_size] for i in range(0, total_books, chunk_size)]
    total_chunks = len(chunks)
    print(f"Total chunks: {total_chunks}")

    # Prepare chunk information with process names
    chunk_infos = [(chunk, f"Process-{i+1}") for i, chunk in enumerate(chunks)]

    with multiprocessing.Pool(processes=num_workers) as pool:
        pool.map(process_chunk, chunk_infos)


def process_chunk(chunk_info):
    chunk_ids, process_name = chunk_info

    chunk_ids = [uuid.UUID(id_) if not isinstance(id_, uuid.UUID) else id_ for id_ in chunk_ids]

    try:
        model = SentenceTransformer('all-MiniLM-L6-v2')
        conn = psycopg2.connect(
        dbname='book_lens',
        user='postgres',
        password='password',
        host='localhost',
        port='5433'
        )

        psycopg2.extras.register_uuid() #need this so uuid shit doesn't break everything

        cursor = conn.cursor(cursor_factory=RealDictCursor)

        # get the data for the chunk
        cursor.execute("""
                       SELECT id, title, book_desc FROM books WHERE id = ANY(%s)
                       """, (chunk_ids,))
        books = cursor.fetchall()

        ids = []
        texts = []
        for book in books:
            book_id = book['id']
            title = book['title'] or ''
            desc = book['book_desc'] or ''
            text = f"{title}. {desc}"
            ids.append(book_id)
            texts.append(text)

        embeddings = model.encode(texts, batch_size=64, show_progress_bar=False)
        embeddings = embeddings / np.linalg.norm(embeddings, axis=1, keepdims=True)

        records = [(str(uuid.uuid4()), str(book_id), embeddings.tolist()) for book_id, embeddings in zip(ids, embeddings)]
        del ids
        del embeddings

        insert_query = """
        INSERT INTO book_embeddings (id, book_id, embedding)
        VALUES (%s, %s, %s);        
        """
        execute_batch(cursor, insert_query, records)
        conn.commit()
        cursor.close()
        conn.close()
        print(f"{process_name}: Processed {len(chunk_ids)} records.")
    except Exception as e:
        print(f"{process_name}: Error processing chunk: {e}", file=sys.stderr)



def get_all_book_ids():
    """
    Gets all of the book ID's from the database, these ID's will be be broken into n chunks for parallell processing
    Each worker will get a chunk of ID's to generate embeddings for it.

    Returns: book_ids, list of all of the book id's
    """
    conn = psycopg2.connect(
        dbname='book_lens',
        user='postgres',
        password='password',
        host='localhost',
        port='5433'
    )
    cursor = conn.cursor()
    cursor.execute("SELECT id FROM books;")
    book_ids = [row[0] for row in cursor.fetchall()] #take the 0th element of each tuple from fetchall and put it in a book_ids list
    cursor.close()
    conn.close()
    return book_ids


if __name__ == "__main__":
    # generate_and_store_embeddings_parallel(num_workers=4, chunk_size=100)
    generate_and_store_embeddings(batch_size=10000)
