function getCoordinatesByRouteId(route_id, callback) {
    fetch("http://" + window.location.host + "/vehicle/status?route_id=" + route_id, {
        method: "GET"
    })
    .then((response) => response.json())
    .then((json) => callback(json.route_id, json.coordinates, json.ts))
}