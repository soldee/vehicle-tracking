package exceptions;

public class VehicleNotFoundException extends VehicleTrackingException {

    public VehicleNotFoundException() {
        super(10, "Vehicle not found");
    }
}
