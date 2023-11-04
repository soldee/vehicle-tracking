import React from "react";
import './FiltersButton.css'


export default function FiltersButton({ onClickHandler }) {

    return (
        <div className="filters-button">
            <button onClick={() => onClickHandler()}>
                <span className="material-symbols-outlined search">search</span>
                FILTERS
            </button>
        </div>
    )
}