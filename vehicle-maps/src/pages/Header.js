import React from "react";
import './Header.css'


export default function Header({ activeTab, tabChangeHandler }) {

    return (
        <div className="tab-buttons">
            <button
                onClick={() => tabChangeHandler(1)}
                className={activeTab === 1 ? 'active' : ''}
            >
                LIVE
                <span className="material-symbols-outlined radio">radio_button_unchecked</span>
            </button>
            <button
                onClick={() => tabChangeHandler(2)}
                className={activeTab === 2 ? 'active' : ''}
            >
                SEARCH
                <span className="material-symbols-outlined filter">filter_alt</span>
            </button>
        </div>
    )
}