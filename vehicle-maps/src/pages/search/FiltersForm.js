import React, { useEffect, useState } from "react";
import './FiltersForm.css'


export default function FiltersForm({ onSearchRouteId, isDataDisplayed }) {

    const [isLoading, setIsLoading] = useState(false)
    const [routeID, setRouteID] = useState("")
    const [userID, setUserID] = useState("")
    const [vehicleID, setVehicleID] = useState("")
    const [dateFrom, setDateFrom] = useState("")
    const [dateTo, setDateTo] = useState("")

    function clearStorage() {
        localStorage.removeItem("route_id")
        localStorage.removeItem("vehicle_id")
        localStorage.removeItem("user_id")
        localStorage.removeItem("date_from")
        localStorage.removeItem("date_to")
    }

    useEffect(() => {
        if (isDataDisplayed) {
            setRouteID(localStorage.getItem("route_id"))
            setUserID(localStorage.getItem("user_id"))
            setVehicleID(localStorage.getItem("vehicle_id"))
            setDateFrom(localStorage.getItem("date_from"))
            setDateTo(localStorage.getItem("date_to"))
        } else {
            clearStorage()
        }
    }, [])

    function handleSubmit(e) {
        e.preventDefault()
        setIsLoading(true)

        clearStorage()

        const params = {}
        if (routeID != null && routeID.trim() != "") {
            params.route_id = routeID
            localStorage.setItem("route_id", routeID);
        }
        if (vehicleID != null && vehicleID.trim() != "") {
            params.vehicle_id = vehicleID
            localStorage.setItem("vehicle_id", vehicleID);
        }
        if (userID != null && userID.trim() != "") {
            params.user_id = userID
            localStorage.setItem("user_id", userID);
        }
        if (dateFrom != null && dateFrom.trim() != "") {
            params['date[gt]'] = dateFrom
            localStorage.setItem("date_from", dateFrom);
        }
        if (dateTo != null && dateTo.trim() != "") {
            params['date[lt]'] = dateTo
            localStorage.setItem("date_to", dateTo);
        }

        fetch(window.REACT_APP_DOMAIN + "/vehicle/status?" + new URLSearchParams(params), { method: "GET" })
            .then(async (response) => {
                setIsLoading(false)
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

    return (
        <form method="GET" onSubmit={handleSubmit} className="filters-form">
            <div className="filters">
                <div className="metadata-filters">
                    <h3>Metadata</h3>
                    <div className="m-filter">
                        <label>
                            Route ID
                            <input value={routeID ?? ""} onChange={(e) => setRouteID(e.target.value)} />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            Vehicle ID
                            <input value={vehicleID ?? ""} onChange={(e) => setVehicleID(e.target.value)} />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            User ID
                            <input value={userID ?? ""} onChange={(e) => setUserID(e.target.value)} />
                        </label>
                    </div>
                </div>
                <div className="date-filters">
                    <h3>Dates</h3>
                    <div className="m-filter">
                        <label>
                            From
                            <input value={dateFrom ?? ""} onChange={(e) => setDateFrom(e.target.value)} />
                        </label>
                    </div>
                    <div className="m-filter">
                        <label>
                            To
                            <input values={dateTo ?? ""} onChange={(e) => setDateTo(e.target.value)} />
                        </label>
                    </div>
                </div>
            </div>
            <button type="submit" className={`submit-button ${isLoading && "button--loading"}`} >
                <span className="button__text">Save</span>
            </button>
        </form>
    )
}