package org.example;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;

import java.util.Random;

public class MqttPublisher {

    public static void main(String[] args) {

        String broker = "tcp://192.168.33.10:1883";
        String username = "guest";
        String password = "guest";
        String clientId = "test-client";

        int QOS = 1;

        MqttClient client;
        try {
            client = new MqttClient(broker, clientId, new MemoryPersistence());
        } catch (MqttException e) {
            throw new RuntimeException(e);
        }

        MqttConnectOptions options = new MqttConnectOptions();
        options.setUserName(username);
        options.setPassword(password.toCharArray());

        try {
            client.connect();
        } catch (MqttException e) {
            throw new RuntimeException(e);
        }

        for (int i=10; i<1000; i++) {
            try {
                String message = "{" +
                        "\"ts\": \"2023-01-01T10:15:" + i + "Z\"," +
                        "\"meta\":{" +
                        "\"vehicle_id\":\"a\"," +
                        "\"user_id\":\"a\"," +
                        "\"route_id\":\"a\"" +
                        "}," +
                        "\"location\":{\"type\":\"Point\",\"coordinates\":[" + new Random().nextInt(10) + "," + new Random().nextInt(10) + "]}," +
                        "\"speed\":" + new Random().nextFloat() * 10 +
                        "}";

                client.publish("test-topic", message.getBytes(), QOS, true);
                Thread.sleep(10000);
            } catch (MqttException e) {
                throw new RuntimeException(e);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
