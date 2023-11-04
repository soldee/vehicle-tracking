import React from 'react';
import { MapContainer, TileLayer } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet'
import RouteList from './RouteList';
import './LeafletMap.css'


delete L.Icon.Default.prototype._getIconUrl;

L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png')
})


function LeafletMap({ data, features }) {

    const center = window.REACT_APP_MAP_CENTER;

    return (
        <div className="map-container">
            <MapContainer center={center} zoom={14} scrollWheelZoom={true} minZoom={13} maxZoom={16}>
                <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />

                <RouteList data={data} features={features} />
            </MapContainer>
        </div>
    );
};

export default LeafletMap;
