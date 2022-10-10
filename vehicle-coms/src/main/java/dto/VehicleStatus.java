package dto;

public class VehicleStatus {

    private String vehicleId;
    private float latitude;
    private float longitude;

    public VehicleStatus(String vehicleId, float latitude, float longitude) {
        this.vehicleId = vehicleId;
        this.latitude = latitude;
        this.longitude = longitude;
    }

    public String getVehicleId() {
        return vehicleId;
    }

    public float getLatitude() {
        return latitude;
    }

    public float getLongitude() {
        return longitude;
    }

    public void setVehicleId(String vehicleId) {
        this.vehicleId = vehicleId;
    }

    public void setLatitude(float latitude) {
        this.latitude = latitude;
    }

    public void setLongitude(float longitude) {
        this.longitude = longitude;
    }
}

