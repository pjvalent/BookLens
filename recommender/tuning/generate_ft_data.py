import pandas as pd
import numpy as np
import json
import psycopg2
import psycopg2.extras
import uuid
import datetime
import re

from typing import List, Dict


def handle_popular_shelves(book) -> List[str]:
    """
    Parse and clean the data that is in the 'popular_shelves' entry of the book
    Should get each of the unique popular shelves that a book has been put in.
    Should filter out any shelves that are 'to read' 'reading' 'currently reading' etc...

    book: the book to process

    returns: 
    """




def generate_dataset(size: int):
    """
    Generate a training dataset to fiine tune the model with the similar books

    size: size of the dataset to produce

    return: n/a
    """
    if size < 1 :
        print("size of generated dataset set to < 1")
        return 0

    

def main():
    status = generate_dataset()


if __name__ == "__main__":
    main()