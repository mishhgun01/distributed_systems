import React, { useState } from "react";

function NoteForm({ note, onSave }) {
    const [title, setTitle] = useState(note ? note.title : "");
    const [content, setContent] = useState(note ? note.content : "");

    const handleSubmit = (e) => {
        e.preventDefault();
        onSave({ id: note?.id, title, content });
    };

    return (
        <form onSubmit={handleSubmit} className="mb-4">
            <div className="mb-3">
                <label htmlFor="title" className="form-label">
                    Title
                </label>
                <input
                    type="text"
                    id="title"
                    className="form-control"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
            </div>
            <div className="mb-3">
                <label htmlFor="content" className="form-label">
                    Content
                </label>
                <textarea
                    id="content"
                    className="form-control"
                    rows="5"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                    required
                ></textarea>
            </div>
            <button type="submit" className="btn btn-primary">
                Save
            </button>
        </form>
    );
}

export default NoteForm;