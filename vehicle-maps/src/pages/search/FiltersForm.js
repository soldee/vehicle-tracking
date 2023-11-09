import React, { useState } from "react";
import './FiltersForm.css'


export default function FiltersForm({ onSearchRouteId }) {

    const [isLoading, setIsLoading] = useState(false)


    function handleSubmit(e) {
        e.preventDefault()
        setIsLoading(true)

        const form = new FormData(e.target)
        const route_id = form.get("route_id_input")
        const vehicle_id = form.get("vehicle_id_input")
        const user_id = form.get("user_id_input")

        fetch(window.REACT_APP_DOMAIN + "/vehicle/status?" + new URLSearchParams({
            route_id: route_id
        }), { method: form.method })
            .then(async (response) => {
                const json = await response.json();
                setIsLoading(false)

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

    return (
        <form method="GET" onSubmit={handleSubmit} className="filters-form">
            <div className="filters">
                <div className="metadata-filters">
                    <h3>Metadata</h3>
                    <div className="m-filter">
                        <label>
                            Route ID
                            <input name="route_id_input" />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            Vehicle ID
                            <input name="vehicle_id_input" />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            User ID
                            <input name="user_id_input" />
                        </label>
                    </div>
                </div>
                <div className="date-filters">
                    <h3>Dates</h3>
                    <div className="m-filter">
                        <label>
                            From
                            <input name="date_from_input" />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            To
                            <input name="date_to_input" />
                        </label>
                    </div>
                </div>
            </div>
            <button type="submit" class={`submit-button ${isLoading && "button--loading"}`} >
                <span class="button__text">Save</span>
            </button>
        </form>
    )
}