package exceptions;

public class VehicleRequestException extends VehicleComsException {

    public VehicleRequestException() {
        super(11, "Vehicle not responding");
    }
}
