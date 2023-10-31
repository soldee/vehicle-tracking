import React from 'react';
import { Marker, Popup, Polyline, FeatureGroup, useMap } from 'react-leaflet';

function Route({vehicle_id, coordinates, route_id, color}) {

    const map = useMap()

    const vehicle_coordinates = coordinates.slice(-1).pop();
    const route_options = {
        weight: 6,
        color: color
    }


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

    return (
        <React.Fragment>
            <FeatureGroup eventHandlers={{ click: onFeatureGroupClick }} >
                <Marker position={vehicle_coordinates} type='vehicle' >
                    <Popup>
                        <dl>
                            <dt><b>Vehicle ID</b> {vehicle_id}</dt>
                            <dt><b>Route ID</b> {route_id}</dt>
                        </dl>
                    </Popup>
                </Marker>
                <Polyline positions={coordinates} pathOptions={route_options} type='route' />
            </FeatureGroup>
        </React.Fragment>
    )

}

export default Route;