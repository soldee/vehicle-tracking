package exceptions;

public class VehicleComsException extends Exception {
    private int code;
    private String message;

    public VehicleComsException(int code, String message) {

    }

    public int getCode() { return code; }

    @Override
    public String getMessage() {
        return message;
    }
}
