from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
from repo.repo import Repo

class MongoRepo(Repo):
    def __init__(self, uri):
        client = connect_to_mongo(uri)
        self.client = client
        self.collection = client["VEHICLE-TRACKING"]["vehicle-status"]

    def insert(self, record):
        self.collection.insert_one(record)

    def close(self):
        print("Closing connection to MongoDB")
        self.client.close()
        
def connect_to_mongo(uri):
    client = MongoClient(uri)
    try:
        client.admin.command('ping')
        print("Connected to MongoDB")
    except Exception as e:
        print("Could not connect to MongoDB. Error is: ", e)
        exit(1)

    return client