# Recommender
This directory serves to hold the recommender service. The backend should use the recommender service to generate recommendations for users based on their books read, ratings, reviews they have written, and other factors not listed here.

### ETL
Extract Transform Load (ETL) serves as both the initial cleaning and ingestion of the data into the system but will also hold the ongoing data engineering work for continuous ingestion of new data into the system. Data can be added via the fe/be services in some cases but the etl process will handle the automated collection and ingestion of large amounts of new data.