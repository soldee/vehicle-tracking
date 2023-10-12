package org.example;

import org.apache.commons.io.IOUtils;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;
import org.eclipse.paho.client.mqttv3.*;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;

import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;
import java.util.UUID;


public class KafkaMqttSource {

    public static void main(String[] args) {

        ClassLoader classloader = Thread.currentThread().getContextClassLoader();
        InputStream is = classloader.getResourceAsStream("application.properties");

        Properties properties = new Properties();
        try {
            properties.load(is);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        String kafkaServer = properties.getProperty("kafka.bootstrap.servers");
        String kafkaTopic = properties.getProperty("kafka.topic");

        String mqttBroker = properties.getProperty("mqtt.broker.uri");
        String mqttUser = properties.getProperty("mqtt.broker.username");
        String mqttPwd = properties.getProperty("mqtt.broker.password");
        String mqttClientId = properties.getProperty("mqtt.client.id");
        String mqttTopic = properties.getProperty("mqtt.topic");
        int QOS = Integer.parseInt(properties.getProperty("mqtt.qos"));


        Properties kafkaProperties = new Properties();
        kafkaProperties.setProperty(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, kafkaServer);
        kafkaProperties.setProperty(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        kafkaProperties.setProperty(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());

        KafkaProducer<String, String> kafkaProducer = new KafkaProducer<>(kafkaProperties);


        final MqttClient client;
        try {
            client = new MqttClient(mqttBroker, UUID.randomUUID().toString(), new MemoryPersistence());

            MqttConnectOptions options = new MqttConnectOptions();
            options.setUserName(mqttUser);
            options.setPassword(mqttPwd.toCharArray());

            client.connect();

        } catch (MqttException e) {
            throw new RuntimeException(e);
        }

        try {
            client.subscribe(mqttTopic, QOS);

            client.setCallback(new MqttCallback() {
                public void connectionLost(Throwable throwable) {
                    System.out.println("ERROR: " + throwable.fillInStackTrace());
                    try {
                        System.out.println("RECONNECTING");
                        client.reconnect();
                        Thread.sleep(100);

                    } catch (MqttException | InterruptedException e) {
                        kafkaProducer.flush();
                        throw new RuntimeException(e);
                    }
                }

                public void messageArrived(String s, MqttMessage mqttMessage) {
                    System.out.println("MQTT MESSAGE (" + s + "): " + mqttMessage.toString());

                    ProducerRecord<String, String> record = new ProducerRecord<>(kafkaTopic, null, mqttMessage.toString());

                    kafkaProducer.send(record, (md, ex) -> {
                        if (ex != null) {
                            System.err.println("exception occurred in producer for review :" + record.value() + ", exception is " + ex);
                            ex.printStackTrace();
                        } else {
                            System.out.println("SEND KAFKA. Partition:" + md.partition() + ", offset:" + md.offset() + ", timestamp:" + md.timestamp());
                        }
                    });
                }

                public void deliveryComplete(IMqttDeliveryToken iMqttDeliveryToken) {
                    try {
                        System.out.println("DELIVERY: " + iMqttDeliveryToken.getMessage().toString());
                    } catch (MqttException e) {
                        throw new RuntimeException(e);
                    }
                }
            });

        } catch (MqttException e) {
            throw new RuntimeException(e);
        } finally {
            try {
                client.close();
            } catch (MqttException ignored) {}
        }


    }
}
