import React, { useEffect } from "react";
import "./styles/DeleteModal.css";

const DeleteModal = ({ onClose, onConfirm }) => {
    useEffect(() => {
        const handleClickOutside = (event) => {
            if (event.target.classList.contains("modal")) {
                onClose();
            }
        };

        document.addEventListener("click", handleClickOutside);
        return () => {
            document.removeEventListener("click", handleClickOutside);
        };
    }, [onClose]);

    return (
        <div className="modal">
            <div className="modal-content">
                <h3>Вы уверены, что хотите удалить эту заметку?</h3>
                <div className="modal-actions">
                    <button className="btn btn-danger" onClick={onConfirm}>
                        Да, удалить
                    </button>
                    <button className="btn btn-secondary" onClick={onClose}>
                        Отмена
                    </button>
                </div>
            </div>
        </div>
    );
};

export default DeleteModal;