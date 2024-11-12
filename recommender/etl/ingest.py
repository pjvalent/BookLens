import pandas as pd
import json
import psycopg2
import psycopg2.extras
import uuid
import datetime
import re
import numpy as np
import os
from dotenv import load_dotenv
from psycopg2.extras import execute_batch, Json
from typing import List, Dict
# from sentence_transformers import SentenceTransformer

load_dotenv()

class GenreHandler:
    def __init__(self):
        """ initialize the genres dictionary """
        self.genres = {}
        with open('../data/goodreads_book_genres_initial.json') as f:
            for line in f:
                obj = json.loads(line)
                self.genres[obj['book_id']] = obj['genres']

    def get_genres(self, book_id):
        genres = self.genres.get(f"{book_id}")
        genres_set = set() #ensure only unique values 
        for key in genres.keys():
            temp = [value.strip() for value in key.split(',')] #format the individual keys
            genres_set.update(temp)
        return {"genres": sorted(genres_set)}

    def _print_genres(self):
        print(self.genres)


def db_connect():
    conn = psycopg2.connect(
        dbname=os.getenv('DB_NAME'),
        user=os.getenv('DB_USER'),
        password=os.getenv('DB_PASSWORD'),
        host=os.getenv('DB_HOST'),
        port=os.getenv('DB_PORT')
    )
    cursor = conn.cursor()

    psycopg2.extras.register_uuid()

    return conn, cursor
    
def normalize_text(text):
     text = text.lower() #lowercase that shiii
     text = re.sub(r'[^a-z0-9\s]', '', text) #remove punctuation
     text = re.sub(r'\s+', ' ', text) #remove extra spaces
     return text.strip()


def ingest_books(batch_size, path):
    books_list = []
    print("ingesting books")
    genre_handler = GenreHandler()
    with open(path, 'r') as f:
        print("opening file")
        for line_number, line in enumerate(f, start=1):
            try:
                book = json.loads(line)
                #TODO: need to filter out any entries that are not english, and where the title doesn't exist in the list already
                #       Reducing the amout that goes to the database is needed

                if book['language_code'] != 'eng':
                     continue #skip this line and proceed to the next one
                if book['country_code'] != 'US':
                     continue #skip this line and proceed to the next one
                if any(x in book['title'].lower() for x in ["box set", "collection", "bundle"]):
                     continue #skip this line and proceed to the next one
                
                #we are now going to push the title/publisher to lowercase, and strip the whitespace from each, then we will check to see if that 
                #book title and publisher combo already exists in the list, get rid of it if it does. And we add a unique constraint to the 
                #table unique (title/publisher) so that we only maintain one copy. If we do find duplicates in our efforts, keep the one with
                #the longer description

                if not any(filter(lambda x: x['title'] == book['title'].lower().strip() 
                                  and x['publisher'] == book['publisher'].lower().strip(), books_list)): #its gross but it works
                    
                    lower_book= {k: v.lower().strip() if k != 'authors' else v for k, v in book.items() if k not in ['popular_shelves', 'series', 'similar_books']}
                    #store the books in lower, ignore the popular_shelves, series, authors, similar_books
                    books_list.append(lower_book)

                else:
                    #ignore the book, don't add it to the book list
                     continue
   
                
            except json.JSONDecodeError as e:
                print(f"Error decoding JSON on line {line_number}: {e}")
                continue  # Skip this line and proceed to the next one


            if line_number % batch_size == 0:
                    # Process the current batch
                    process_books(books_list, genre_handler)
                    # Clear the list for the next batch
                    books_list = []

            if line_number % (batch_size * 2) == 0:
                    print(f"Processed {line_number} lines...")

        if books_list:
            process_books(books_list, genre_handler)


def process_books(books_list, genre_handler):
    conn, cursor = db_connect()

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
        author_id = parse_author_info(book.get('authors'))
        genres = genre_handler.get_genres(book_id=book['book_id'])
        genres_json = Json(genres)

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
              book_format,
              author_id,
              genres_json
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
            format,
            author_id,
            genres
        )
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        ON CONFLICT DO NOTHING;
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


def ingest_genres(path: str):
    """
    Pull the genre info for each book based on id. The genres are a k:v pair
    Where k is the book id and v is dict of generes and their count {genre: count}
    there can be multiple generes in a single genere:count pair
    We want to 
    """
    genres_source = []
    genres_set = set()

    with open(path) as f:
         for line in f:
              genres_source.append(json.loads(line))
    
    for book in genres_source:
        dict_of_generes = book.get('genres') #this is a dict of genere: num_occurance where the genere can be none, one, or many genres
        for key in dict_of_generes.keys():
            genres = [genre.strip() for genre in key.split(',')]
            genres_set.update(genres) #Update the set with the union of the passed in set


    conn, cursor = db_connect()

    records = []
    for genre in genres_set:
        id = uuid.uuid4()
        genere_name = genre

        record = (
              id,
              genere_name
         )
        records.append(record)

        insert_query = """
        INSERT INTO generes (
            genere_id,
            genere_name
        )
        VALUES (%s, %s)
        ON CONFLICT DO NOTHING;
    """

    # Execute batch insert
    execute_batch(cursor, insert_query, records)

    # Commit changes and close connection
    conn.commit()
    cursor.close()
    conn.close()


def parse_author_info(authors: List[Dict[str, int]]) -> int:
    """
    Parse the first author form the list of authors of the book. 

    authors: list of authors, each entry being {"author": '00000', "role": "{role}"} 
                TODO: taking just the first author for now but we want to eventually modify the schema to support multiple authors
    """
    if not authors:
        return 0

    first_author = authors[0]
    return first_author['author_id']


def ingest_authors(path):
    """
    Ingest the authors and populate the author table

    Path: the path to the authors json file

    Returns: n/a
    """

    author_data = []

    #since the authors file is only ~100mb we just read the whole thing in
    #each entry of the author data list is a json object
    with open(path) as f:
        for line in f:
            author_data.append(json.loads(line))

    conn, cursor = db_connect()

    records = []
    for author in author_data:
         id = uuid.uuid4()
         author_name = author.get('name')
         average_rating = author.get('average_rating')
         author_id = author.get('author_id')
         text_review_count = author.get('text_reviews_count')
         ratings_count = author.get('ratings_count')

         record = (
              id,
              author_name,
              average_rating,
              author_id,
              text_review_count,
              ratings_count
         )

         records.append(record)
    
    insert_query = """
        INSERT INTO authors (
            id,
            author_name,
            average_rating,
            author_id,
            text_review_count,
            ratings_count
        )
        VALUES (%s, %s, %s, %s, %s, %s);
    """

    # Execute batch insert
    execute_batch(cursor, insert_query, records)

    # Commit changes and close connection
    conn.commit()
    cursor.close()
    conn.close()


def main():
    # genre_handler = GenreHandler()
    # genres = genre_handler.get_genres(29360997)
    # print(genres)
    #'29360997': {'fantasy, paranormal': 158, 'mystery, thriller, crime': 53, 'fiction': 23}

    # ingest_genres('../data/goodreads_book_genres_initial.json')
    print("Calling ingest books")
    ingest_books(5000, '../data/goodreads_books.json')
    # print("calling ingest authors")
    # ingest_authors('../data/goodreads_book_authors.json')



if __name__ == "__main__":
    main()