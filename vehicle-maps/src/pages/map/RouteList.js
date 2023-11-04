import React from "react";
import Route from "./Route";

export default function RouteList({ data, features }) {

    return (
        <React.Fragment>
            {
                data.map(vehicle => {
                    return <Route 
                        key={vehicle.vehicle_id} 
                        coordinates={vehicle.coordinates}
                        timestamps={vehicle.ts}
                        speeds={vehicle.speed}
                        vehicle_id={vehicle.vehicle_id} 
                        route_id={vehicle.route_id}
                        color="#0fff"
                        route_points_toggled={features.route_points_toggled}
                        focus_on_click={features.focus_on_click}
                    />
                })
            }
        </React.Fragment>
    )
}