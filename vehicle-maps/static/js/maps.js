const map = L.map('map', {
    minZoom: 13,
    maxZoom: 16
})
map.setView([41.396561, 2.159583], 14);

const tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
	attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);



var route_coordinates = [[41.395573, 2.155684], [41.396013, 2.157112], [41.396491, 2.159460]];
var route = L.polyline(route_coordinates, {
    weight: 6,
    color: "#0fff"
});
route.properties = {
    type: "route"
}

var vehicle = L.marker(route_coordinates.slice(-1).pop());
vehicle.bindPopup("vehicle_id: 1")
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

