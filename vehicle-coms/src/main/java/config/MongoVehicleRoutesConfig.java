package config;

import com.mongodb.MongoClientSettings;
import com.mongodb.MongoCredential;
import com.mongodb.ServerAddress;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import model.RouteModel;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.MongoDatabaseFactory;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.SimpleMongoClientDatabaseFactory;

import java.util.Collections;


@Configuration
public class MongoVehicleRoutesConfig {

    @Value("mongo.vehicleActions.db")
    private String db;

    @Value("mongo.vehicleActions.host")
    private String host;

    @Value("mongo.vehicleActions.port")
    private int port;

    @Value("mongo.vehicleActions.username")
    private String username;

    @Value("mongo.vehicleActions.password")
    private String password;

    @Value("mongo.vehicleActions.collection")
    private String collection;



    @Bean(name = "vehicleRoutesMongoClient")
    public MongoClient mongoClient() {

        MongoCredential credential = MongoCredential
                .createCredential(username, db, password.toCharArray());

        return MongoClients.create(MongoClientSettings.builder()
                .applyToClusterSettings(builder -> builder
                        .hosts(Collections.singletonList(new ServerAddress(host,port))))
                .credential(credential)
                .build());
    }

    @Bean(name = "vehicleRoutesMongoCollection")
    public MongoCollection<RouteModel> mongoCollection(
            @Qualifier("vehiclesMongoClient") MongoClient mongoClient) {
        return mongoClient.getDatabase(db).getCollection(collection, RouteModel.class);
    }

    @Bean(name = "vehicleRoutesMongoFactory")
    public MongoDatabaseFactory mongoDatabaseFactory(
            @Qualifier("vehicleActionsMongoClient") MongoClient mongoClient) {
        return new SimpleMongoClientDatabaseFactory(mongoClient, db);
    }

    @Bean(name = "vehicleRoutesMongoTemplate")
    public MongoTemplate mongoTemplate(@Qualifier("vehicleRoutesMongoFactory") MongoDatabaseFactory mongoDatabaseFactory) {
        return new MongoTemplate(mongoDatabaseFactory);
    }

}
