package dto;

import org.json.JSONArray;
import org.json.JSONObject;

import java.util.ArrayList;
import java.util.Arrays;

public class VehicleAllStatusDto {

    private ArrayList<VehicleStatusDto> vehicleStatusDtos;


    public VehicleAllStatusDto(ArrayList<VehicleStatusDto> vehicleStatusDtos) {
        this.vehicleStatusDtos = vehicleStatusDtos;
    }

    public String toJSONString() {
        return new JSONObject()
                .append("status", new JSONArray(vehicleStatusDtos))
                .toString();
    }

    public ArrayList<VehicleStatusDto> getVehicleStatuses() {
        return vehicleStatusDtos;
    }

    public void setVehicleStatuses(ArrayList<VehicleStatusDto> vehicleStatusDtos) {
        this.vehicleStatusDtos = vehicleStatusDtos;
    }
}
