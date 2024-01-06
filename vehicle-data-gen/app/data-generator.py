import threading
import os
import csv
import time
import random
import datetime
import uuid
import sys
from repo.mqtt_repo import MqttRepo
from repo.mongo_repo import MongoRepo
from repo.repo import Repo
from dotenv import load_dotenv

def runner(repo, csv_path):
    print("Reading " + csv_path)

    route_id = str(uuid.uuid4())
    user_id = os.path.basename(csv_path).split(".")[0]
    vehicle_id = user_id

    with open(csv_path) as csv_file:
        reader = csv.reader(csv_file, delimiter=",")
        for row in reader:
            lat = float(row[0])
            long = float(row[1])
            speed = random.uniform(0.1, 4)
            print("{},{},{}".format(lat,long,speed))

            repo.insert(record=generate_status(route_id, user_id, vehicle_id, speed, lat, long))
            time.sleep(5)


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


if __name__ == "__main__":

    if len(sys.argv) != 2:
        print("Expected 1 argument")
        exit(1)

    load_dotenv()

    match sys.argv[1]:
        case "--mongo":
            mongo_uri = os.getenv("MONGO_URI")
            if mongo_uri == "":
                print("MONGO_URI env variable is required to connect to MongoDB")
                exit(1)
            repo = MongoRepo(mongo_uri)
        case "--mqtt":
            repo = MqttRepo()
        case _:
            print(f"Invalid argument provided for repo {sys.argv[0]}. Valid arguments are '--mongo' or '--mqtt'")
            exit(1)

    scriptDirPath = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
    dataDir = os.path.join(scriptDirPath, "data")
    csvList = os.listdir(dataDir)

    threadPool = []

    for i in range(len(csvList)):
        csv_path = os.path.join(dataDir, csvList[i])
        t = threading.Thread(target=runner, args=(repo,csv_path))
        threadPool.append(t)
        t.start()
        time.sleep(random.uniform(0.1,3))

    for thread in threadPool:
        thread.join()

