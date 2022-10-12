package services;

import com.mongodb.client.MongoCollection;
import constants.VehicleConstants;
import exceptions.VehicleNotFoundException;
import exceptions.VehicleRequestException;
import exceptions.VehicleStartException;
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

@Service
public class VehicleControlService {

    @Autowired
    @Qualifier("vehiclesMongoCollection")
    private MongoCollection<VehicleModel> mongoVehiclesCollection;

    @Autowired
    @Qualifier("vehiclesMongoCollection")
    private MongoCollection<VehicleModel> mongoActionsCollection;

    @Autowired
    private VehicleStatusService statusService;


    public boolean startVehicle(String vehicleId) throws VehicleNotFoundException, VehicleRequestException {
        VehicleModel vehicle = statusService.getVehicle(vehicleId);
        if(vehicle == null) throw new VehicleNotFoundException();

        return vehicleStartRequest(vehicle.getURL());
    }

    public boolean stopVehicle(String vehicleId) throws VehicleNotFoundException, VehicleRequestException {
        VehicleModel vehicle = statusService.getVehicle(vehicleId);
        if(vehicle == null) throw new VehicleNotFoundException();

        return vehicleStopRequest(vehicle.getURL());
    }


    private boolean vehicleStartRequest(String url) throws VehicleRequestException {
        try {
            String statusUrl = url + "/start";
            HttpGet post = new HttpGet(statusUrl);
            HttpClient client = HttpClientBuilder.create().build();

            HttpResponse response = client.execute(post);
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


    private boolean vehicleStopRequest(String url) throws VehicleRequestException {
        try {
            String statusUrl = url + "/stop";
            HttpGet post = new HttpGet(statusUrl);
            HttpClient client = HttpClientBuilder.create().build();

            HttpResponse response = client.execute(post);
            InputStream responseStream = response.getEntity().getContent();

            String jsonString = IOUtils.toString(responseStream, StandardCharsets.UTF_8);
            JSONObject json = new JSONObject(jsonString);

            String vehicleResponse = json.getString(VehicleConstants.STOP);
            if (vehicleResponse.equals("OK")) return true;
            else throw new VehicleStartException();
        } catch (Exception e) {
            throw new VehicleRequestException();
        }
    }

}
