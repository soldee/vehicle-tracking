package exceptions;

public class VehicleStartException extends VehicleComsException{

    public VehicleStartException() { super(12, "Vehicle not started"); }
}
