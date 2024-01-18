import threading
import os
import csv
import time
import random
import uuid
import sys
from repo.mqtt_repo import MqttRepo
from repo.mongo_repo import MongoRepo
from dotenv import load_dotenv

run = True

def runner(repo, csv_path):
    print("Reading " + csv_path)

    route_id = str(uuid.uuid4())
    user_id = os.path.basename(csv_path).split(".")[0]
    vehicle_id = user_id

    with open(csv_path) as csv_file:
        reader = csv.reader(csv_file, delimiter=",")
        global run
        for row in reader:
            if not run:
                repo.close()
                return
            lat = float(row[0])
            long = float(row[1])
            speed = random.uniform(0.1, 4)
            print("{},{},{}".format(lat,long,speed))

            repo.insert_status(route_id, user_id, vehicle_id, speed, lat, long)
            time.sleep(5)


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
            mqtt_broker = os.getenv("MQTT_BROKER")
            mqtt_port = os.getenv("MQTT_PORT")
            mqtt_username = os.getenv("MQTT_USERNAME")
            mqtt_password = os.getenv("MQTT_PASSWORD")
            mqtt_client_id = os.getenv("MQTT_CLIENT_ID")
            if mqtt_broker == "" or mqtt_port == "" or mqtt_username == "" or mqtt_password == "" or mqtt_client_id == "":
                print("MQTT_BROKER, MQTT_PORT, MQTT_USERNAME, MQTT_PASSWORD and MQTT_CLIENT_ID env variables are required to connect to MQTT broker")
                exit(1)
            
            repo = MqttRepo(mqtt_broker, int(mqtt_port), mqtt_username, mqtt_password, mqtt_client_id)
        case _:
            print(f"Invalid argument provided for repo {sys.argv[0]}. Valid arguments are '--mongo' or '--mqtt'")
            exit(1)

    scriptDirPath = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
    dataDir = os.path.join(scriptDirPath, "data")
    csvList = os.listdir(dataDir)

    threadPool = []

    try:
        for i in range(len(csvList)):
            csv_path = os.path.join(dataDir, csvList[i])
            t = threading.Thread(target=runner, args=(repo,csv_path))
            threadPool.append(t)
            t.start()
            time.sleep(random.uniform(0.1,3))

        while True:
            for thread in threadPool:
                if not thread.is_alive():
                    threadPool.remove(thread)
                else:
                    time.sleep(.5)
    except KeyboardInterrupt:
        print("Keyboard interrupt received. Shutting down threads")
        run = False