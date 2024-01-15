from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
from repo.repo import Repo
from datetime import datetime

class MongoRepo(Repo):
    def __init__(self, uri):
        client = connect_to_mongo(uri)
        self.client = client
        self.collection = client["VEHICLE-TRACKING"]["vehicle-status"]

    def insert_status(self, route_id, user_id, vehicle_id, speed, lat, long):
        self.collection.insert_one(generate_status(route_id, user_id, vehicle_id, speed, lat, long))

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

def generate_status(route_id, user_id, vehicle_id, speed, latitude, longitude):
    return {
        'ts': datetime.datetime.now(), 
        'meta': {
            'route_id': route_id, 
            'user_id': user_id, 
            'vehicle_id': vehicle_id
        }, 
        'speed': speed, 
        'location': {
            'type': 'Point', 
            'coordinates': [latitude, longitude]
        }
    }