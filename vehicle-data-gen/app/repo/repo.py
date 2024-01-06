from abc import ABC, abstractmethod

class Repo(ABC):
    @abstractmethod
    def insert(self, record):
        pass