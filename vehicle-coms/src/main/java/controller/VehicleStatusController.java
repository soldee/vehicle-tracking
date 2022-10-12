package controller;

import constants.VehicleConstants;
import dto.VehicleAllStatusDto;
import dto.VehicleStatusDto;
import exceptions.VehicleNotFoundException;
import exceptions.VehicleComsException;
import model.VehicleModel;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import services.VehicleStatusService;


@RestController
public class VehicleStatusController {

    @Autowired
    private VehicleStatusService vehicleStatusService;

    @GetMapping(path = "/info", produces = MediaType.APPLICATION_JSON_VALUE)
    public String info(@RequestParam String vehicleId) {
        try {
            VehicleModel vehicle = vehicleStatusService.getVehicle(vehicleId);
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, vehicle.toJSONString())
                    .toString();

        } catch (Exception e) {
            int errorCode = VehicleConstants.UNKNOWN_ERROR_CODE;
            if (e instanceof VehicleComsException) errorCode = ((VehicleComsException) e).getCode();
            return new JSONObject()
                    .append(VehicleConstants.CODE, errorCode)
                    .append(VehicleConstants.ERROR_MESSAGE, e.getMessage())
                    .toString();
        }
    }


    @GetMapping(path = "/status", produces = MediaType.APPLICATION_JSON_VALUE)
    public String status(@RequestParam String vehicleId) {
        try {
            VehicleStatusDto status = vehicleStatusService.getStatus(vehicleId);
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, status)
                    .toString();

        } catch (Exception e) {
            int errorCode = VehicleConstants.UNKNOWN_ERROR_CODE;
            if (e instanceof VehicleComsException) errorCode = ((VehicleComsException) e).getCode();
            return new JSONObject()
                    .append(VehicleConstants.CODE, errorCode)
                    .append(VehicleConstants.ERROR_MESSAGE, e.getMessage())
                    .toString();
        }
    }

    
    @GetMapping(path = "/status/all", produces = MediaType.APPLICATION_JSON_VALUE)
    public String allStatus() {
        try {
            VehicleAllStatusDto status = vehicleStatusService.getAllstatus();
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, status)
                    .toString();

        } catch (Exception e) {
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.UNKNOWN_ERROR_CODE)
                    .append(VehicleConstants.ERROR_MESSAGE, e.getMessage())
                    .toString();
        }
    }



}
