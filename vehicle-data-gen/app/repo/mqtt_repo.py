import json
import time
from repo.repo import Repo
from paho.mqtt import client as mqtt_client
from datetime import date, datetime

class MqttRepo(Repo):
    def __init__(self, broker, port, username, password, client_id):
        client = mqtt_client.Client(client_id)
        client.username_pw_set(username, password)
        client.on_connect = on_connect
        client.connect(broker, port)
        self.client = client

        while not client.is_connected:
            time.sleep(.2)

    def insert_status(self, route_id, user_id, vehicle_id, speed, lat, long):
        json_str = json.dumps(generate_status(route_id, user_id, vehicle_id, speed, lat, long), default=serializer)
        self.client.publish("test-topic", json_str)

    def close(self):
        self.client.disconnect()


def on_connect(client, userdata, flags, rc):
    if rc == 0:
        print("Connected to MQTT broker")
    else:
        print("Failed to connect. rc={}".format(rc))
        exit(1)


def serializer(obj):
    if isinstance(obj, (datetime, date)):
        return obj.isoformat()
    raise TypeError ("type %s not serializable" % type(obj))


def generate_status(route_id, user_id, vehicle_id, speed, lat, long):
    return {
        'ts': datetime.now().strftime("%Y-%m-%dT%H:%M:%SZ"), 
        'meta': {
            'route_id': route_id, 
            'user_id': user_id, 
            'vehicle_id': vehicle_id
        }, 
        'speed': speed, 
        'location': [lat, long]
    }
