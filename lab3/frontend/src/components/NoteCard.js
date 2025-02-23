import React, { useState } from "react";
import "./styles/NoteList.css";

const NoteCard = ({ note, onEdit, onDelete }) => {
    const [showOptions, setShowOptions] = useState(false);

    return (
        <div className="note-card" onClick={() => setShowOptions(!showOptions)}>
            <h3>{note.title}</h3>
            <p>{note.content}</p>
            {showOptions && (
                <div className="note-options">
                    <button className="btn btn-primary" onClick={() => onEdit(note)}>
                        Редактировать
                    </button>
                    <button className="btn btn-danger" onClick={() => onDelete(note.id)}>
                        Удалить
                    </button>
                </div>
            )}
        </div>
    );
};

export default NoteCard;