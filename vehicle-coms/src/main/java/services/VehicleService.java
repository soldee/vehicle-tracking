package services;

import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoCursor;
import constants.VehicleConstants;
import dto.VehicleAllStatusDto;
import dto.VehicleStatusDto;
import exceptions.VehicleNotFoundException;
import model.VehicleModel;
import org.apache.commons.io.IOUtils;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.HttpClientBuilder;
import org.bson.Document;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;

@Service
public class VehicleService {

    @Autowired
    @Qualifier("vehiclesMongoCollection")
    private MongoCollection<VehicleModel> mongoVehiclesCollection;

    @Autowired
    @Qualifier("vehiclesMongoCollection")
    private MongoCollection<VehicleModel> mongoActionsCollection;


    public VehicleModel getVehicle(String vehicleId) throws VehicleNotFoundException {
        VehicleModel vehicle = mongoVehiclesCollection.find(Document.parse("{_id:"+vehicleId+"}")).first();
        if (vehicle == null) throw new VehicleNotFoundException();
        return vehicle;
    }


    public VehicleStatusDto getStatus(String vehicleId) throws VehicleNotFoundException {
        VehicleModel vehicle = mongoVehiclesCollection.find(Document.parse("{_id:"+vehicleId+"}")).first();
        if (vehicle == null) throw new VehicleNotFoundException();
        if (!vehicle.isActive()) return new VehicleStatusDto(vehicle.getObjectId().toString(), vehicle.getLatitude(), vehicle.getLongitude(), vehicle.getLastTimestamp().toString());

        return vehicleGetStatusRequest(vehicle.getURL());
    }


    public VehicleAllStatusDto getAllstatus() {
        ArrayList<VehicleStatusDto> statusDtos = new ArrayList<>();

        MongoCursor<VehicleModel> it = mongoVehiclesCollection.find(Document.parse("{}")).cursor();
        while (it.hasNext()) {
            VehicleModel vehicle = it.next();
            statusDtos.add(new VehicleStatusDto(vehicle.getObjectId().toString(), vehicle.getLatitude(), vehicle.getLongitude(), vehicle.getLastTimestamp().toString()));
        }
        return new VehicleAllStatusDto(statusDtos);
    }


    private VehicleStatusDto vehicleGetStatusRequest(String url) {
        try {
            String statusUrl = url + "/status";
            HttpGet post = new HttpGet(statusUrl);
            HttpClient client = HttpClientBuilder.create().build();

            HttpResponse response = client.execute(post);
            InputStream responseStream = response.getEntity().getContent();

            String jsonString = IOUtils.toString(responseStream, StandardCharsets.UTF_8);
            JSONObject json = new JSONObject(jsonString);

            String vehicleId = json.getString(VehicleConstants.VEHICLE_ID);
            float latitude = json.getFloat(VehicleConstants.LATITUDE);
            float longitude = json.getFloat(VehicleConstants.LONGITUDE);
            String timestamp = json.getString(VehicleConstants.TIMESTAMP);

            return new VehicleStatusDto(vehicleId, latitude, longitude, timestamp);
        } catch (Exception e) {
            return null;
        }
    }

}
