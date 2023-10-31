import React, { useState } from 'react';
import './App.css';
import LeafletMap from './pages/map/LeafletMap';
import RouteForm from './pages/map/RouteForm';

export default function App() { 

    const [vehiclesData, setVehiclesData] = useState([])

    function searchRouteHandler(json) {
        const data = [json]
        setVehiclesData(data)
    }

    return (
        <div className="App">
        <LeafletMap data={vehiclesData} />
        <RouteForm onSearchRouteId={searchRouteHandler} />
        </div>
    );
}