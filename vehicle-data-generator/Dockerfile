FROM ubuntu:23.10

RUN apt-get update && apt-get install -y \ 
    python3-pip \
    python3.11-venv \
    && apt-get clean

WORKDIR /home/vehicle-data-generator
COPY ./requeriments.txt requeriments.txt

RUN python3 -m venv . --system-site-packages
RUN ./bin/pip install -r requeriments.txt

COPY ./app src/app/
COPY ./data src/data/

ENV MONGO_URI=mongodb://root:root@mongo:27017/?authSource=admin&replicaSet=rs0
ENV MQTT_BROKER=192.168.33.10
ENV MQTT_PORT=1883
ENV MQTT_USERNAME=guest
ENV MQTT_PASSWORD=guest
ENV MQTT_CLIENT_ID=test-client

ENTRYPOINT ["./bin/python", "-u", "src/app/data-generator.py"]
CMD ["--mqtt"]