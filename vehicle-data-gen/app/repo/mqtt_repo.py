from repo.repo import Repo

class MqttRepo(Repo):
    def __init__(self):
        print("connect to mqtt broker")

    def insert(self, record):
        print("publish record to mqtt broker: ", record)