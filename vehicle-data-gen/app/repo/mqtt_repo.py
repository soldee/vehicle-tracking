from repo.repo import Repo
from paho.mqtt import client as mqtt_client

class MqttRepo(Repo):
    def __init__(self):
        print("connect to mqtt broker")

    def insert(self, record):
        print("publish record to mqtt broker: ", record)

    def close(self):
        print("closing connection to MQTT broker")