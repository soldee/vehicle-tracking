import React, { useState } from 'react';
import './App.css';
import Header from './pages/Header';
import TabContent from './pages/TabContent';

export default function App() { 

    const [activeTab, setActiveTab] = useState(1)

    return (
        <div className="app">
            <Header activeTab={activeTab} tabChangeHandler={(tabNumber) => setActiveTab(tabNumber)} />
            <TabContent activeTab={activeTab} />
        </div>
    );
}