package exceptions;

public class VehicleStopException extends VehicleComsException{

    public VehicleStopException() { super(12, "Vehicle failed to stop"); }
}
