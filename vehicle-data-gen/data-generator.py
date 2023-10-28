from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
import threading
import os
import csv
import time
import random
import datetime
import uuid


def connect_to_mongo():
    client = MongoClient("mongodb://root:root@192.168.33.10:27017/?authSource=admin&replicaSet=rs0")
                
    try:
        client.admin.command('ping')
        print("Connected to MongoDB!")
    except Exception as e:
        print(e)
        exit(1)

    return client


def runner(collection, csvPath):
    print("Reading " + csvPath)

    route_id = str(uuid.uuid4())
    user_id = os.path.basename(csv_path).split(".")[0]
    vehicle_id = user_id
    speed = random.uniform(0.1, 4)

    with open(csvPath) as csv_file:
        reader = csv.reader(csv_file, delimiter=",")
        for row in reader:
            lat = float(row[0])
            long = float(row[1])
            print("{},{}".format(lat,long))

            collection.insert_one(generateStatus(route_id, user_id, vehicle_id, speed, lat, long))
            time.sleep(5)


def generateStatus(route_id, user_id, vehicle_id, speed, latitude, longitude):
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




if __name__ == "__main__":
    DB = "VEHICLE-TRACKING"
    COLLECTION = "vehicle-status"

    client = connect_to_mongo()
    collection = client["VEHICLE-TRACKING"]["vehicle-status"]

    scriptDirPath = os.path.dirname(os.path.realpath(__file__))
    dataDir = os.path.join(scriptDirPath, "data")
    csvList = os.listdir(dataDir)

    threadPool = []

    for i in range(len(csvList)):
        csv_path = os.path.join(dataDir, csvList[i])
        t = threading.Thread(target=runner, args=(collection,csv_path))
        threadPool.append(t)
        t.start()
        time.sleep(random.uniform(0.1,3))

    for thread in threadPool:
        thread.join()

