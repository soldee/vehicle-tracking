import React, { useState } from "react";
import LeafletMap from "../map/LeafletMap";
import MapFeatures from "../map/MapFeatures";


export default function LiveTab({ mapFeatures, onMapFeaturesChange }) {

    const vehiclesData = []

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