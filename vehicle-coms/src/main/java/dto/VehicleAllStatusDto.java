package dto;

import java.util.ArrayList;

public class VehicleAllStatusDto {

    private ArrayList<VehicleStatusDto> vehicleStatusDtos;


    public VehicleAllStatusDto(ArrayList<VehicleStatusDto> vehicleStatusDtos) {
        this.vehicleStatusDtos = vehicleStatusDtos;
    }

    public ArrayList<VehicleStatusDto> getVehicleStatuses() {
        return vehicleStatusDtos;
    }

    public void setVehicleStatuses(ArrayList<VehicleStatusDto> vehicleStatusDtos) {
        this.vehicleStatusDtos = vehicleStatusDtos;
    }
}
