import React, { useState } from 'react';
import { Marker, Popup, Polyline, FeatureGroup, useMap } from 'react-leaflet';
import marker from '../../assets/black-route-point.png'
import L from 'leaflet'


function Route({vehicle_id, coordinates, timestamps, speeds, route_id, route_points_toggled, focus_on_click}) {

    const map = useMap()
    const [color, _] = useState(colors[(Math.floor(Math.random() * colors.length))])

    const route_color = color[0]
    const vehicle_marker_color = color[1]

    const vehicle_coordinates = coordinates.slice(-1).pop();
    const route_options = {
        weight: 8,
        color: route_color
    }

    const pointsIcon = new L.Icon({
        iconUrl: marker,
        iconRetinaUrl: marker,
        iconSize: window.ROUTE_POINTS_ICON_SIZE
    })

    const finalIcon = new L.DivIcon({
        html: `<span class="material-symbols-outlined" style="color: ${vehicle_marker_color}; font-variation-settings: 'FILL' 1, 'wght' 200, 'GRAD' 0, 'opsz' 24; font-size:40px">location_on</span>`,
        iconAnchor: [20,36],
        className: 'd',
        popupAnchor: [0,-20],
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
                <Marker position={vehicle_coordinates} type='vehicle' icon={finalIcon} >
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
                                <dl>
                                    <dt>
                                        <b>Timestamp</b> {
                                            new Date(timestamps[index]).toLocaleString("es-ES", {timeZone: "Europe/Madrid"})
                                        }
                                    </dt>
                                    <dt>
                                        <b>Speed</b> {
                                            speeds[index]
                                        } m/s
                                    </dt>
                                </dl>
                            </Popup>
                        </Marker>
                    })
                }
            </FeatureGroup>
        </React.Fragment>
    )

}

const colors = [
    ["#FEC101", "#CB9A00"],
    ["#03ED27", "#02BD1F"],
    ["#A15DEA", "#804ABB"],
    ["#C9C0FF", "#A099CC"],
    ["#94E5DF", "#76B7B2"],
    ["#A6CCD3", "#84A3A8"],
    ["#BB0059", "#950047"],
    ["#088DA5", "#067084"],
    ["#D99B7D", "#AD7C64"],
    ["#F5ECDF", "#C4BCB2"]
]

export default Route;