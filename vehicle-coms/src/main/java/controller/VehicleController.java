package controller;

import constants.VehicleConstants;
import dto.VehicleAllStatusDto;
import dto.VehicleStatusDto;
import exceptions.VehicleTrackingException;
import model.VehicleModel;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import services.VehicleService;


@RestController
public class VehicleController {

    @Autowired
    private VehicleService vehicleService;

    @PostMapping(path = "/info",
            consumes = MediaType.APPLICATION_JSON_VALUE,
            produces = MediaType.APPLICATION_JSON_VALUE)
    public String info(@RequestParam String vehicleId) {
        try {
            VehicleModel vehicle = vehicleService.getVehicle(vehicleId);
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, vehicle.toJSONString())
                    .toString();

        } catch (VehicleTrackingException e) {
            return new JSONObject()
                    .append(VehicleConstants.CODE, e.getCode())
                    .append(VehicleConstants.ERROR_MESSAGE, e.getMessage())
                    .toString();
        } catch (Exception e) {
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.UNKNOWN_ERROR_CODE)
                    .append(VehicleConstants.ERROR_MESSAGE, e.getMessage())
                    .toString();
        }
    }


    @GetMapping(path = "/status", produces = MediaType.APPLICATION_JSON_VALUE)
    public String status(@RequestParam String vehicleId) {
        try {
            VehicleStatusDto status = vehicleService.getStatus(vehicleId);
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

    
    @GetMapping(path = "/status/all", produces = MediaType.APPLICATION_JSON_VALUE)
    public String allStatus() {
        try {
            VehicleAllStatusDto status = vehicleService.getAllstatus();
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
