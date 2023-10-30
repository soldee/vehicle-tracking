import React from 'react';
import './App.css';
import LeafletMap from './pages/map/LeafletMap';
import RouteForm from './pages/map/RouteForm';

export default function App() { 

    function searchRouteHandler(data) {
        const coordinates = data.coordinates;
        const route_id = data.route_id;

        console.log(coordinates, route_id)
    }

    return (
        <div className="App">
        <LeafletMap />
        <RouteForm onSearchRouteId={searchRouteHandler} />
        </div>
    );
}