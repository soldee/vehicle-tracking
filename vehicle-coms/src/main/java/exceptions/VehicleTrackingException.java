package exceptions;

public class VehicleTrackingException extends Exception {
    private int code;
    private String message;

    public VehicleTrackingException(int code, String message) {

    }

    public int getCode() { return code; }

    @Override
    public String getMessage() {
        return message;
    }
}
