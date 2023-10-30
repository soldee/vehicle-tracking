import React from 'react';
import { Marker, Popup, Polyline, FeatureGroup, useMap } from 'react-leaflet';

function Route({coordinates, route_id, color}) {

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
                        Route Id: {route_id}
                    </Popup>
                </Marker>
                <Polyline positions={coordinates} pathOptions={route_options} type='route' />
            </FeatureGroup>
        </React.Fragment>
    )

}

export default Route;