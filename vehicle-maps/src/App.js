import React, { useState } from 'react';
import './App.css';
import LeafletMap from './pages/map/LeafletMap';
import RouteForm from './pages/map/RouteForm';

export default function App() { 

    const [vehiclesData, setVehiclesData] = useState([])
    const [mapFeatures, setMapFeatures] = useState({
        route_points_toggled: false
    })

    function searchRouteHandler(json) {
        const data = [json]
        setVehiclesData(data)
    }

    function mapFeatureChangeHandler(features) {
        setMapFeatures((prev) => {
            return {
                ...prev,
                features
            }
        })
    }

    return (
        <div className="App">
            <LeafletMap 
                data={vehiclesData} 
                features={mapFeatures} 
            />
            <RouteForm 
                onSearchRouteId={searchRouteHandler} 
                mapFeatures={mapFeatures} 
                onMapFeaturesChange={mapFeatureChangeHandler} 
            />
        </div>
    );
}