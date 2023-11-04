import React from 'react';
import { Marker, Popup, Polyline, FeatureGroup, useMap } from 'react-leaflet';
import marker from '../../assets/black-route-point.png'
import L from 'leaflet'


function Route({vehicle_id, coordinates, timestamps, route_id, color, route_points_toggled, focus_on_click}) {

    const map = useMap()

    const vehicle_coordinates = coordinates.slice(-1).pop();
    const route_options = {
        weight: 6,
        color: color
    }

    const pointsIcon = new L.Icon({
        iconUrl: marker,
        iconRetinaUrl: marker,
        iconSize: window.ROUTE_POINTS_ICON_SIZE
    })


    function onFeatureGroupClick(e) {
        Object.values(e.target._layers).forEach(layer => {
            switch (layer.options.type) {
                case "vehicle":
                    layer.openPopup()
                    break;
                case "route":
                    if (focus_on_click) {
                        map.fitBounds(layer.getBounds())
                    }
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
                {
                    route_points_toggled &&
                    coordinates.map((c, index) => {
                        return <Marker key={c} position={c} icon={pointsIcon} eventHandlers={{
                            mouseover: (e) => e.target.openPopup(),
                            mouseout: (e) => e.target.closePopup()
                        }}>
                            <Popup>
                                <b>Timestamp</b> {
                                    new Date(timestamps[index]).toLocaleDateString("es-ES", { 
                                        weekday: 'short', year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', second:'2-digit'
                                    })
                                }
                            </Popup>
                        </Marker>
                    })
                }
            </FeatureGroup>
        </React.Fragment>
    )

}

export default Route;