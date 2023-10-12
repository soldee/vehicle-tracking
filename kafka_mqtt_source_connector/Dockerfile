FROM openjdk:8-jdk-alpine

WORKDIR /app
COPY ./target/kafka_mqtt_source_connector.jar kafka_mqtt_source_connector.jar

CMD ["java", "-jar", "kafka_mqtt_source_connector.jar"]