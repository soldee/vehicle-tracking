package dto;

import org.json.JSONObject;

public class VehicleStatusDto {

    private String vehicleId;
    private Float latitude;
    private Float longitude;
    private String timestamp;
    private boolean active;

    public VehicleStatusDto(String vehicleId, Float latitude, Float longitude, String timestamp) {
        this.vehicleId = vehicleId;
        this.latitude = latitude;
        this.longitude = longitude;
        this.timestamp = timestamp;
    }

    public VehicleStatusDto(String vehicleId, Float latitude, Float longitude, String timestamp, boolean active) {
        this.vehicleId = vehicleId;
        this.latitude = latitude;
        this.longitude = longitude;
        this.timestamp = timestamp;
        this.active = active;
    }

    public String toJSONString() {
        return new JSONObject()
                .append("vehicle_id", vehicleId)
                .append("latitude", latitude)
                .append("longitude", longitude)
                .append("timestamp", timestamp)
                .append("active", active)
                .toString();
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

    public void setLatitude(Float latitude) {
        this.latitude = latitude;
    }

    public void setLongitude(Float longitude) {
        this.longitude = longitude;
    }

    public String getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(String timestamp) {
        this.timestamp = timestamp;
    }
}

