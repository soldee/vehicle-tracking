package services;

import com.mongodb.client.MongoCollection;
import constants.VehicleConstants;
import exceptions.VehicleNotFoundException;
import exceptions.VehicleRequestException;
import exceptions.VehicleStartException;
import exceptions.VehicleStopException;
import model.RouteModel;
import model.VehicleModel;
import org.apache.commons.io.IOUtils;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.HttpClientBuilder;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.Date;

@Service
public class VehicleControlService {

    @Autowired
    private VehicleStatusService statusService;

    @Autowired
    @Qualifier("vehicleRoutesMongoCollection")
    private MongoCollection<RouteModel> vehicleRoutesCollection;


    public boolean startVehicle(String vehicleId) throws VehicleNotFoundException, VehicleRequestException {
        VehicleModel vehicle = statusService.getVehicle(vehicleId);
        if(vehicle == null) throw new VehicleNotFoundException();

        boolean hasStarted = vehicleStartRequest(vehicle.getURL(), vehicle.getObjectId().toString());
        if (hasStarted) vehicleRoutesCollection.insertOne(
                new RouteModel(VehicleConstants.START, vehicleId, new Date(System.currentTimeMillis())));
        return hasStarted;
    }

    public boolean stopVehicle(String vehicleId) throws VehicleNotFoundException, VehicleRequestException {
        VehicleModel vehicle = statusService.getVehicle(vehicleId);
        if(vehicle == null) throw new VehicleNotFoundException();
        
        String route_id = vehicleStopRequest(vehicle.getURL());
        vehicleRoutesCollection.insertOne(
                new RouteModel(VehicleConstants.STOP, vehicleId, new Date(System.currentTimeMillis()), route_id));
        return true;
    }


    private boolean vehicleStartRequest(String url, String route_id) throws VehicleRequestException {
        try {
            String statusUrl = url + "/start";
            HttpGet get = new HttpGet(statusUrl);
            get.addHeader(VehicleConstants.ROUTE_ID, route_id);
            HttpClient client = HttpClientBuilder.create().build();

            HttpResponse response = client.execute(get);
            InputStream responseStream = response.getEntity().getContent();

            String jsonString = IOUtils.toString(responseStream, StandardCharsets.UTF_8);
            JSONObject json = new JSONObject(jsonString);

            String vehicleResponse = json.getString(VehicleConstants.START);
            if (vehicleResponse.equals("OK")) return true;
            else throw new VehicleStartException();
        } catch (Exception e) {
            throw new VehicleRequestException();
        }
    }


    private String vehicleStopRequest(String url) throws VehicleRequestException {
        try {
            String statusUrl = url + "/stop";
            HttpGet post = new HttpGet(statusUrl);
            HttpClient client = HttpClientBuilder.create().build();

            HttpResponse response = client.execute(post);
            InputStream responseStream = response.getEntity().getContent();

            String jsonString = IOUtils.toString(responseStream, StandardCharsets.UTF_8);
            JSONObject json = new JSONObject(jsonString);

            String vehicleResponse = json.getString(VehicleConstants.STOP);
            if (vehicleResponse.equals("OK")) return json.getString(VehicleConstants.ROUTE_ID);
            else throw new VehicleStopException();
        } catch (Exception e) {
            throw new VehicleRequestException();
        }
    }

}
