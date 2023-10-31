import React, { useEffect, useState } from "react";
import Route from "./Route";

export default function MapFeatures({ data }) {

    const [isLoading, setIsLoading] = useState(true);
    const [vehicle, setVehicle] = useState(null);

    /*useEffect(() => {
        fetch(window.REACT_APP_DOMAIN + "/vehicle/status?" + new URLSearchParams({
            route_id: "1ae5540f-3a95-42e3-8d6b-5bd1f804d3e3"
        }))
            .then((response) => {
                if (!response.ok) {
                    throw new Error(response.status)
                } 
                return response.json()
            })
            .then((json) => {
                setVehicle({coordinates: json.coordinates, route_id: json.route_id});
                setIsLoading(false);
            })
            .catch(err => {
                alert(err);
            })
    }, [window.REACT_APP_DOMAIN])*/

    return (
        <React.Fragment>
            {
                data.map(vehicle => {
                    return <Route 
                        key={vehicle.vehicle_id} 
                        coordinates={vehicle.coordinates} 
                        vehicle_id={vehicle.vehicle_id} 
                        route_id={vehicle.route_id} 
                        color="#0fff" 
                    />
                })
            }
        </React.Fragment>
    )
}