import psycopg2
from psycopg2.extras import RealDictCursor, execute_batch
import numpy as np
from sentence_transformers import SentenceTransformer
from tqdm import tqdm
import uuid

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

if __name__ == "__main__":
    generate_and_store_embeddings(batch_size=1000)
