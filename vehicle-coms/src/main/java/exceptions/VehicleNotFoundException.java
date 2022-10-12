package exceptions;

public class VehicleNotFoundException extends VehicleComsException {

    public VehicleNotFoundException() {
        super(10, "Vehicle not found");
    }
}
