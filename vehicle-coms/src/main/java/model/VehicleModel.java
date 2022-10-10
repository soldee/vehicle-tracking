package model;

import org.bson.types.ObjectId;
import org.json.JSONObject;
import org.springframework.data.mongodb.core.mapping.MongoId;

import java.util.Date;

public class VehicleModel {

    @MongoId
    private ObjectId objectId;
    private String vehicleName;
    private boolean active;
    private float latitude;
    private float longitude;
    private Date lastTimestamp;

    public VehicleModel(ObjectId objectId, String vehicleName, boolean active, float latitude, float longitude, Date lastTimestamp) {
        this.objectId = objectId;
        this.vehicleName = vehicleName;
        this.active = active;
        this.latitude = latitude;
        this.longitude = longitude;
        this.lastTimestamp = lastTimestamp;
    }

    public String toJSONString() {
        return new JSONObject()
                .append("vehicle_id", objectId)
                .append("vehicle_name", vehicleName)
                .append("active", active)
                .append("latitude", latitude)
                .append("longitude", longitude)
                .toString();
    }

    public ObjectId getObjectId() {
        return objectId;
    }

    public String getVehicleName() {
        return vehicleName;
    }

    public boolean isActive() {
        return active;
    }

    public void setObjectId(ObjectId objectId) {
        this.objectId = objectId;
    }

    public void setVehicleName(String vehicleName) {
        this.vehicleName = vehicleName;
    }

    public void setActive(boolean active) {
        this.active = active;
    }

    public float getLatitude() {
        return latitude;
    }

    public void setLatitude(float latitude) {
        this.latitude = latitude;
    }

    public float getLongitude() {
        return longitude;
    }

    public void setLongitude(float longitude) {
        this.longitude = longitude;
    }

    public Date getLastTimestamp() {
        return lastTimestamp;
    }

    public void setLastTimestamp(Date lastTimestamp) {
        this.lastTimestamp = lastTimestamp;
    }
}
