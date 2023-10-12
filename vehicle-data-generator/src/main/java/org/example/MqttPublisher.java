package org.example;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;

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

        for (int i=0; i<1000; i++) {
            try {
                client.publish("test-topic", ("message" + i).getBytes(), QOS, true);
                Thread.sleep(3000);
            } catch (MqttException e) {
                throw new RuntimeException(e);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
