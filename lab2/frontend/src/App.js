import React, { useState, useEffect } from "react";
import NoteList from "./components/NoteList";
import NoteModal from "./components/NoteModal";
import DeleteModal from "./components/DeleteModal";

function App() {
    const [notes, setNotes] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [noteToEdit, setNoteToEdit] = useState(null);
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [noteToDelete, setNoteToDelete] = useState(null);

    useEffect(() => {
        fetchNotes();
    }, []);

    const fetchNotes = () => {
        fetch("http://localhost:8080/api/notes")
            .then((response) => response.json())
            .then((data) => setNotes(data))
            .catch((error) => console.error("Error fetching notes:", error));
    };

    const handleAddNote = (note) => {
        fetch("http://localhost:8080/api/notes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(note),
        })
            .then((response) => response.json())
            .then((savedNote) => {
                setNotes([...notes, savedNote]);
                setShowModal(false);
            })
            .catch((error) => console.error("Error saving note:", error));
    };

    const handleEditNote = (note) => {
        fetch(`http://localhost:8080/api/notes`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(note),
        })
            .then((response) => response.json())
            .then((updatedNote) => {
                setNotes(
                    notes.map((n) => (n.id === updatedNote.id ? updatedNote : n))
                );
                setShowModal(false);
                setNoteToEdit(null);
            })
            .catch((error) => console.error("Error updating note:", error));
    };

    const handleDeleteNote = () => {
        fetch(`http://localhost:8080/api/notes/${noteToDelete}`, {
            method: "DELETE",
        })
            .then((response) => {
                if (response.ok) {
                    setNotes(notes.filter((note) => note.id !== noteToDelete));
                    setShowDeleteModal(false);
                    setNoteToDelete(null);
                }
            })
            .catch((error) => console.error("Error deleting note:", error));
    };

    return (
        <div className="container mt-5">
            <div className="d-flex justify-content-between align-items-center mb-4">
                <h1>Заметки</h1>
                <button
                    className="btn btn-primary"
                    onClick={() => {
                        setNoteToEdit(null);
                        setShowModal(true);
                    }}
                >
                    Добавить заметку
                </button>
            </div>

            <NoteList
                notes={notes}
                onEdit={(note) => {
                    setNoteToEdit(note);
                    setShowModal(true);
                }}
                onDelete={(id) => {
                    setNoteToDelete(id);
                    setShowDeleteModal(true);
                }}
            />

            {showModal && (
                <NoteModal
                    note={noteToEdit}
                    onClose={() => setShowModal(false)}
                    onSave={noteToEdit ? handleEditNote : handleAddNote}
                />
            )}

            {showDeleteModal && (
                <DeleteModal
                    onClose={() => setShowDeleteModal(false)}
                    onConfirm={handleDeleteNote}
                />
            )}
        </div>
    );
}

export default App;