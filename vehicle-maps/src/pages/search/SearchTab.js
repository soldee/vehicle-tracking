import React, { useState } from "react";
import MapFeatures from '../map/MapFeatures'
import LeafletMap from "../map/LeafletMap";
import Filters from "./Filters";


export default function SearchTab({ mapFeatures, onMapFeaturesChange}) {
    
    const [vehiclesData, setVehiclesData] = useState([])
    
    function searchRouteHandler(json) {
        const data = [json]
        setVehiclesData(data)
    }

    return (
        <React.Fragment>
            <MapFeatures mapFeatures={mapFeatures} onMapFeaturesChange={onMapFeaturesChange} />
            <Filters onSearchRouteId={searchRouteHandler} />
            <LeafletMap
                data={vehiclesData}
                features={mapFeatures}
            />
        </React.Fragment>
    )
}