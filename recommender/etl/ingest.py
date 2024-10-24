import pandas as pd
import json
import psycopg2
import psycopg2.extras
import uuid
import datetime
from psycopg2.extras import execute_batch
# from sentence_transformers import SentenceTransformer



def ingest_books(batch_size, path):
    books_list = []
    print("ingesting books")

    with open(path, 'r') as f:
        print("opening file")
        for line_number, line in enumerate(f, start=1):
            print("Line: ", line_number)
            try:
                book = json.loads(line)
                
                books_list.append(book)
                
            except json.JSONDecodeError as e:
                print(f"Error decoding JSON on line {line_number}: {e}")
                continue  # Skip this line and proceed to the next one


            if line_number % batch_size == 0:
                    # Process the current batch
                    process_books(books_list)
                    # Clear the list for the next batch
                    books_list = []

            if line_number % (batch_size * 2) == 0:
                    print(f"Processed {line_number} lines...")

        if books_list:
            process_books(books_list)


def process_books(books_list):
    conn = psycopg2.connect(
        dbname='',
        user='',
        password='',
        host='',
        port=''
    )
    cursor = conn.cursor()

    psycopg2.extras.register_uuid()

    records = []
    for book in books_list:
        id = uuid.uuid4()
        author = str(uuid.uuid4())
        price = 0
        created_at = datetime.datetime.now()
        updated_at = datetime.datetime.now()
        isbn = to_int_zero(book.get('isbn13'))
        title = book.get('title')
        num_pages = to_int_zero(book.get('num_pages'))
        publication_day = to_int_zero(book.get('publication_day'))
        publication_month = to_int_zero(book.get('publication_month'))
        publication_year = to_int_zero(book.get('publication_year'))
        publisher = book.get('publisher')
        book_desc = book.get('description')
        book_format = book.get('format')


        record = (
              id,
              author,
              price,
              created_at,
              updated_at,
              isbn,
              title,
              num_pages,
              publication_day,
              publication_month,
              publication_year,
              publisher,
              book_desc,
              book_format
         )
        records.append(record)

        insert_query = """
        INSERT INTO books (
            id,
            author,
            price,
            created_at,
            updated_at,
            isbn,
            title,
            num_pages,
            publication_day,
            publication_month,
            publication_year,
            publisher,
            book_desc,
            format
        )
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        ON CONFLICT (isbn) DO UPDATE SET
            author = EXCLUDED.author,
            price = EXCLUDED.price,
            title = EXCLUDED.title,
            created_at = EXCLUDED.created_at,
            updated_at = EXCLUDED.updated_at,
            num_pages = EXCLUDED.num_pages,
            publication_day = EXCLUDED.publication_day,
            publication_month = EXCLUDED.publication_month,
            publication_year = EXCLUDED.publication_year,
            publisher = EXCLUDED.publisher,
            book_desc = EXCLUDED.book_desc,
            format = EXCLUDED.format;
    """

    # Execute batch insert
    execute_batch(cursor, insert_query, records)

    # Commit changes and close connection
    conn.commit()
    cursor.close()
    conn.close()


def to_int_zero(value):
     try:
          return int(value)
     except (TypeError, ValueError):
          return 0


def main():
    print("Calling ingest books")
    ingest_books(1000, '../data/goodreads_books.json')


if __name__ == "__main__":
    main()