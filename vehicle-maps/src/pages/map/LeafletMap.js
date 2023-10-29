import React from 'react';
import { MapContainer, TileLayer } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet'
import Coordinates from './Coordinates';

delete L.Icon.Default.prototype._getIconUrl;

L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png')
})


function LeafletMap() {

    const center = [41.389260,2.131592];

    return (
        <div className="map-container">
            <MapContainer center={center} zoom={14} scrollWheelZoom={true} minZoom={13} maxZoom={16}>
                <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />

                <Coordinates />
            </MapContainer>
        </div>
    );
};

export default LeafletMap;
