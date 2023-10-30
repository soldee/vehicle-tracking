import React, { useEffect, useState } from "react";
import Route from "./Route";

export default function MapFeatures() {

    const [isLoading, setIsLoading] = useState(true);
    const [vehicle, setVehicle] = useState(null);

    useEffect(() => {
        fetch(window.REACT_APP_DOMAIN + "/vehicle/status?" + new URLSearchParams({
            route_id: "22c472d8-ba14-4295-a5c9-972a4691ed41"
        }))
            .then((response) => response.json())
            .then((json) => {
                setVehicle({coordinates: json.coordinates, route_id: json.route_id});
                setIsLoading(false);
            })
            .catch(err => {
                alert(err);
            })
    }, [])

    return (
        <React.Fragment>
            { !isLoading && 
                <Route 
                    coordinates={vehicle.coordinates} 
                    route_id={vehicle.route_id} 
                    color="#0fff" 
                /> 
            }
            <Route coordinates={[[41.388577,2.128103], [41.389357,2.130848], [41.389560,2.131592]]} vehicle_id={2} color="#000f" />
        </React.Fragment>
    )
}