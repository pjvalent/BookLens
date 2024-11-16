import pandas as pd
import numpy as np
import json
import psycopg2
import psycopg2.extras
import uuid
import datetime
import re

from typing import List, Dict


def handle_popular_shelves(shelves: List[Dict[str, str]]) -> List[str]:
    """
    Parse and clean the data that is in the 'popular_shelves' entry of the book
    Should get each of the unique popular shelves that a book has been put in.
    Should filter out any shelves that are 'to read' 'reading' 'currently reading' etc...

    shelves: list of k:v pairs, usually {'count':'#', 'name':'{shelve}'}

    returns: a list of all the unique shelves that a book belongs to
    """

    unique_shelves = set() #sets ensure uniqueness when adding elements

    list_of_bad_shelves = ['read', 'school', 'grade', 'chick-lit', '2009-books', 'chicklit', 'owned', 'borrowed', 'library', 'hardcover', 'books-i-own', 'on-my-shelf', 'on-the-shelf']

    for entry in shelves:
        if int(entry.get('count')) < 1:
            # book is on a shelf less than one time which is not possible, dont add shelf to the set
            continue
        shelf = entry.get('name').lower().strip()
        if any(word in shelf for word in list_of_bad_shelves):
            continue
        unique_shelves.add(entry.get('name').lower().strip()) #lowercase and strip the whitespace

    return sorted(unique_shelves)



def generate_dataset(size: int, path: str):
    """
    Generate a training dataset to fiine tune the model with the similar books

    size: size of the dataset to produce

    path: the path to the file location of the books json

    return: 0 or 1: 0 if fail 1 if success
    """
    if size < 1 :
        print("size of generated dataset set to < 1")
        return 0
    
    books_list = []
    with open(path, 'r') as f:
        for line_number, line in enumerate(f, start=1):
            try:
                book = json.loads(line)
                shelves = book.get('popular_shelves')
                unique_shelves = handle_popular_shelves(shelves)
                print(unique_shelves)
            except Exception as e:
                print("exception occured: ", e)
                return 0

            if line_number > 5:
                return 1

    

def main():
    status = generate_dataset(10, '../data/goodreads_books.json')


if __name__ == "__main__":
    main()