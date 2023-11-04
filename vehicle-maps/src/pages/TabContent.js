import React, { useState } from "react";
import LiveTab from "./live/LiveTab";
import SearchTab from "./search/SearchTab";


export default function TabContent({ activeTab }) {

    const [mapFeatures, setMapFeatures] = useState({
        route_points_toggled: false,
        focus_on_click: false
    })

    function mapFeatureChangeHandler(features) {
        setMapFeatures({
            route_points_toggled: features.route_points_toggled,
            focus_on_click: features.focus_on_click
        })
    }


    return (
        <div className="tab-content">
            {activeTab === 1 && (
                <LiveTab mapFeatures={mapFeatures} onMapFeaturesChange={mapFeatureChangeHandler} />
            )}
            {activeTab === 2 && (
                <SearchTab mapFeatures={mapFeatures} onMapFeaturesChange={mapFeatureChangeHandler} />
            )}
        </div>
    )
}