package services;

import com.mongodb.client.MongoCollection;
import dto.VehicleStatus;
import exceptions.VehicleNotFoundException;
import model.VehicleModel;
import org.bson.Document;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

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

    public VehicleStatus getStatus(String vehicleId) throws VehicleNotFoundException {
        VehicleModel vehicle = mongoVehiclesCollection.find(Document.parse("{_id:"+vehicleId+"}")).first();
        if (vehicle == null) throw new VehicleNotFoundException();
        if (!vehicle.isActive()) return new VehicleStatus(vehicle.getObjectId().toString(), vehicle.getLatitude(), vehicle.getLongitude());

        // TODO call to vehicle to retrieve information
        return null;
    }
}
