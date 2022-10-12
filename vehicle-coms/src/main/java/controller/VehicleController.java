package controller;

import constants.VehicleConstants;
import exceptions.VehicleComsException;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import services.VehicleControlService;


@RestController
public class VehicleController {

    @Autowired
    private VehicleControlService vehicleControlService;


    @GetMapping(path = "/start", produces = MediaType.APPLICATION_JSON_VALUE)
    public String start(@RequestParam String vehicleId) {
        try {
            boolean start = vehicleControlService.startVehicle(vehicleId);
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, start)
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


    @GetMapping(path = "/stop", produces = MediaType.APPLICATION_JSON_VALUE)
    public String stop(@RequestParam String vehicleId) {
        try {
            boolean start = vehicleControlService.stopVehicle(vehicleId);
            return new JSONObject()
                    .append(VehicleConstants.CODE, VehicleConstants.OK_CODE)
                    .append(VehicleConstants.RESPONSE, start)
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
}
