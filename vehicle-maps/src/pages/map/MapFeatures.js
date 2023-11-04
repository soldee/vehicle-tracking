import React from "react";
import './MapFeatures.css'


export default function MapFeatures({ mapFeatures, onMapFeaturesChange }) {

    function handleMapFeatureChange(type, change) {
        mapFeatures[type] = change
        onMapFeaturesChange(mapFeatures)
    }

    return (
        <div className="features-container">
            <label>
                <p>Toggle route points</p>
                <input 
                    type="checkbox"
                    checked={mapFeatures.route_points_toggled}
                    onChange={(e) => handleMapFeatureChange("route_points_toggled", e.target.checked)} 
                />
            </label>
            <label>
                <p>Focus on click</p>
                <input 
                    type="checkbox"
                    checked={mapFeatures.focus_on_click}
                    onChange={(e) => handleMapFeatureChange("focus_on_click", e.target.checked)}
                />
            </label>
        </div>
    );
}