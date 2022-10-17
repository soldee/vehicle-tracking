package model;

import java.util.Date;

public class RouteModel {

    private String action;
    private String vehicleId;
    private Date timestamp;
    private String route_id;
    private String user_id;

    public RouteModel(String action, String vehicleId, Date timestamp, String user_id) {
        this.action = action;
        this.vehicleId = vehicleId;
        this.timestamp = timestamp;
    }

    public RouteModel(String action, String vehicleId, Date timestamp, String user_id, String route_id) {
        this.action = action;
        this.vehicleId = vehicleId;
        this.timestamp = timestamp;
        this.route_id = route_id;
    }

    public Date getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(Date timestamp) {
        this.timestamp = timestamp;
    }

    public String getVehicleId() {
        return vehicleId;
    }

    public void setVehicleId(String vehicleId) {
        this.vehicleId = vehicleId;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }

    public String getRoute_id() {
        return route_id;
    }

    public void setRoute_id(String route_id) {
        this.route_id = route_id;
    }
}
