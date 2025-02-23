import React, { useState, useEffect } from "react";
import "./styles/NoteModal.css";

const NoteModal = ({ note, onClose, onSave }) => {
    const [title, setTitle] = useState(note ? note.title : "");
    const [content, setContent] = useState(note ? note.content : "");

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

    const handleSave = () => {
        onSave({ id: note?.id, title, content });
        onClose();
    };

    return (
        <div className="modal">
            <div className="modal-content">
                <h3>{note ? "Редактировать заметку" : "Новая заметка"}</h3>
                <input
                    type="text"
                    placeholder="Заголовок"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                <textarea
                    placeholder="Содержание"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                ></textarea>
                <div className="modal-actions">
                    <button className="btn btn-primary" onClick={handleSave}>
                        Сохранить
                    </button>
                    <button className="btn btn-secondary" onClick={onClose}>
                        Отмена
                    </button>
                </div>
            </div>
        </div>
    );
};

export default NoteModal;