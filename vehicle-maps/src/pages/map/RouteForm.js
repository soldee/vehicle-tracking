import React from "react";

export default function RouteForm({ onSearchRouteId, mapFeatures, onMapFeaturesChange }) {

    function handleSubmit(e) {
        e.preventDefault()

        const form = new FormData(e.target)
        const route_id = form.get("route_id_input")

        fetch(window.REACT_APP_DOMAIN + "/vehicle/status?" + new URLSearchParams({
            route_id: route_id
        }), { method: form.method })
            .then(async (response) => {
                const json = await response.json();

                if (!response.ok) {
                    const err = json.error
                    throw new Error(response.status + " " + err)
                }
                else {
                    onSearchRouteId(json)
                }
            })
            .catch(err => {
                alert(err)
            })
    }


    function handleMapFeatureChange(type, change) {
        switch(type){
            case "route_points_toggled":
                mapFeatures[type] = change
                onMapFeaturesChange(mapFeatures)
                break;
            default:
                break;
        }
    }


    return (
        <div className="features-form-container">
            <form method="GET" onSubmit={handleSubmit}>
                <label>
                    Route ID: 
                    <input name="route_id_input" />
                </label>
                <button type="submit">Search</button>
            </form>
            <label>
                Toggle route points
                <input 
                    type="checkbox"
                    checked={mapFeatures.route_points_toggled}
                    onChange={(e) => handleMapFeatureChange("route_points_toggled", e.target.checked)} 
                />
            </label>
        </div>
    );
}