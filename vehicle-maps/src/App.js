import React, { useState } from 'react';
import './App.css';
import LeafletMap from './pages/map/LeafletMap';
import RouteForm from './pages/map/RouteForm';

export default function App() { 

    const [activeTab, setActiveTab] = useState(1)

    const [vehiclesData, setVehiclesData] = useState([])
    const [mapFeatures, setMapFeatures] = useState({
        route_points_toggled: false
    })

    function searchRouteHandler(json) {
        const data = [json]
        setVehiclesData(data)
    }

    function mapFeatureChangeHandler(features) {
        setMapFeatures((prev) => {
            return {
                ...prev,
                features
            }
        })
    }

    return (
        <div className="app">
            <div className="tab-buttons">
                <button
                    onClick={() => setActiveTab(1)}
                    className={activeTab === 1 ? 'active' : ''}
                >
                    LIVE
                    <span className="material-symbols-outlined radio">radio_button_unchecked</span>
                </button>
                <button
                    onClick={() => setActiveTab(2)}
                    className={activeTab === 2 ? 'active' : ''}
                >
                    SEARCH
                    <span className="material-symbols-outlined filter">filter_alt</span>
                </button>
            </div>
            <div className="tab-content">
                {activeTab === 1 && (
                    <React.Fragment>
                        <LeafletMap 
                            data={vehiclesData} 
                            features={mapFeatures} 
                        />
                    </React.Fragment>
                )}
                {activeTab === 2 && (
                    <React.Fragment>
                        <RouteForm onSearchRouteId={searchRouteHandler} mapFeatures={mapFeatures} onMapFeaturesChange={mapFeatureChangeHandler} />
                        <LeafletMap 
                            data={vehiclesData} 
                            features={mapFeatures} 
                        />
                    </React.Fragment>
                )}
            </div>
        </div>
    );
}