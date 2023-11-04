import React from "react";
import './Modal.css';
import { createPortal } from 'react-dom'


export default function Modal(props) {

    return (
        <React.Fragment>
            {createPortal(
                <div className="modal-container">
                    <div className="modal">
                        <button className="modal-button" onClick={() => props.onClose()}>
                            <span className="material-symbols-outlined close">close</span>
                        </button>
                        {props.children}
                    </div>
                </div>,
                document.getElementById("modal")
            )}
        </React.Fragment>
    )
}