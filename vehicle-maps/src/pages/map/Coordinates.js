import React from 'react';
import { Marker, Popup, Polyline, FeatureGroup, useMap } from 'react-leaflet';

function Coordinates() {

    const map = useMap()


    function onFeatureGroupClick(e) {
        Object.values(e.target._layers).forEach(layer => {
            switch (layer.options.type) {
                case "vehicle":
                    layer.openPopup()
                    break;
                case "route":
                    map.fitBounds(layer.getBounds())
                    break;
            }
        })
    }


    function renderCoordinates() {
        const route_coordinates = [[41.388477,2.128103], [41.389157,2.130848], [41.389260,2.131592]]
        const vehicle_coordinates = route_coordinates.slice(-1).pop();
        const route_options = {
            weight: 6,
            color: "#0fff"
        }

        return (
            <FeatureGroup eventHandlers={{ click: onFeatureGroupClick }} >
                <Marker position={vehicle_coordinates} type='vehicle' >
                    <Popup>
                        A pretty CSS3 popup. <br /> Easily customizable.
                    </Popup>
                </Marker>
                <Polyline positions={route_coordinates} pathOptions={route_options} type='route' />
            </FeatureGroup>
        )
    }

    return (
        <div>
            {renderCoordinates()}
        </div>
    )

}

export default Coordinates;