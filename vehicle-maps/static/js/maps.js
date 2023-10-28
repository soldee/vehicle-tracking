const map = L.map('map', {
    minZoom: 13,
    maxZoom: 16
})
map.setView([41.396561, 2.159583], 14);

const tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
	attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

getCoordinatesByRouteId("19ba0b8f-d10d-4c29-a47f-4d796e1d1242", renderCoordinates)


function renderCoordinates(route_id, route_coordinates, route_timestamps) {

    var route = L.polyline(route_coordinates, {
        weight: 6,
        color: "#0fff"
    });
    route.properties = {
        type: "route"
    }
    
    var vehicle = L.marker(route_coordinates.slice(-1).pop());
    vehicle.bindPopup(`vehicle_id: 1\nroute_id: ${route_id}`)
    vehicle.properties = {
        type: "vehicle",
        vehicle_id: 1
    }
    
    var route_group = L.featureGroup([route, vehicle]).addTo(map)
    route_group.on('click', function(e) {
    
        route_group.getLayers().forEach(layer => {
            switch (layer.properties.type) {
                case "vehicle":
                    console.log(layer.properties.type + " - " + layer.properties.vehicle_id)
                    layer.openPopup()
                    break;
                case "route":
                    console.log(layer.properties.type)
                    map.fitBounds(layer.getBounds());
                    break;
            }
        })
    });
}
