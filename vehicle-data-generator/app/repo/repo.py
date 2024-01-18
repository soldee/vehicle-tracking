from abc import ABC, abstractmethod

class Repo(ABC):
    @abstractmethod
    def insert_status(self, route_id, user_id, vehicle_id, speed, lat, long):
        pass

    @abstractmethod
    def close(self):
        pass