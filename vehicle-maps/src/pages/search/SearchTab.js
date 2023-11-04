import React, { useState } from "react";
import MapFeatures from '../map/MapFeatures'
import LeafletMap from "../map/LeafletMap";
import FiltersForm from "./FiltersForm";
import './SearchTab.css'
import FiltersButton from "./FiltersButton";
import Modal from "../../UI/Modal/Modal";


export default function SearchTab({ mapFeatures, onMapFeaturesChange}) {
    
    const [vehiclesData, setVehiclesData] = useState([])
    const [showFilters, setShowFilters] = useState(false)
    
    function searchRouteHandler(json) {
        const data = [json]
        setVehiclesData(data)
        setShowFilters(false)
    }

    return (
        <React.Fragment>
            <div className="map-header">
                <MapFeatures mapFeatures={mapFeatures} onMapFeaturesChange={onMapFeaturesChange} />
                <FiltersButton onClickHandler={() => setShowFilters(true)} />
            </div>
            { showFilters && 
                <Modal onClose={() => setShowFilters(false)} > 
                    <FiltersForm onSearchRouteId={searchRouteHandler} />
                </Modal>    
            }
            <LeafletMap
                data={vehiclesData}
                features={mapFeatures}
            />
        </React.Fragment>
    )
}