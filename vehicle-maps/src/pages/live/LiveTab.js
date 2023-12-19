import React, { useEffect, useState } from "react";
import LeafletMap from "../map/LeafletMap";
import MapFeatures from "../map/MapFeatures";


export default function LiveTab({ mapFeatures, onMapFeaturesChange }) {

    const [vehiclesData, setVehiclesData] = useState([])

    useEffect(() => {
        const sse = new EventSource("http://localhost:8000/vehicle/status/subscribe")

        sse.onmessage = (e) => {
            console.log("Received generic message: " + e.data)
        }
        sse.onerror = (e) => {
            console.log("received error: ", e)
        }

        sse.addEventListener("status", (e) => {
            const newVehicleStatus = JSON.parse(e.data)
            setVehiclesData((prevData => {
                const i = prevData.findIndex(vehicleData => vehicleData.route_id == newVehicleStatus.route_id)
                if (i != -1) {
                    const newData = [...prevData]
                    const newStatus = {...newData[i]}
                    const newTs = [...newStatus.ts, newVehicleStatus.ts]
                    const newCoordinates = [...newStatus.coordinates, newVehicleStatus.coordinates]
                    const newSpeed = [...newStatus.speed, newVehicleStatus.speed]
                    newStatus.ts = newTs
                    newStatus.coordinates = newCoordinates
                    newStatus.speed = newSpeed
                    newData[i] = newStatus
                    return newData
                } else {
                    return [...prevData, {
                        "ts": [newVehicleStatus.ts],
                        "coordinates": [newVehicleStatus.coordinates],
                        "speed": [newVehicleStatus.speed],
                        "route_id": newVehicleStatus.route_id,
                        "vehicle_id": newVehicleStatus.vehicle_id,
                        "user_id": newVehicleStatus.user_id
                    }]
                }
            }))
        })

        sse.addEventListener("close", (e) => {
            console.log("Received close event from origin: " + e.data)
            sse.close()
        })

        return () => {
            sse.close()
        }
    }, [])

    return (
        <React.Fragment>
            <MapFeatures mapFeatures={mapFeatures} onMapFeaturesChange={onMapFeaturesChange} />
            <LeafletMap
                data={vehiclesData}
                features={mapFeatures}
            />
        </React.Fragment>
    )
}