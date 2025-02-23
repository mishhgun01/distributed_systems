import React, { useState, useEffect } from "react";
import axios from "axios";
import NoteList from "./components/NoteList";
import NoteModal from "./components/NoteModal";
import SearchBar from "./components/SearchBar";
import "./App.css";

const App = () => {
    const [notes, setNotes] = useState([]);
    const [filteredNotes, setFilteredNotes] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [currentNote, setCurrentNote] = useState(null);

    const API_BASE_URL = "http://localhost:8080/api/notes";

    useEffect(() => {
        const fetchNotes = async () => {
            try {
                const response = await axios.get(API_BASE_URL);
                setNotes(response.data);
                setFilteredNotes(response.data);
            } catch (error) {
                console.error("Ошибка при загрузке заметок:", error);
            }
        };

        fetchNotes();
    }, []);

    const handleSearch = (query) => {
        if (!query) {
            setFilteredNotes(notes);
        } else {
            const lowerCaseQuery = query.toLowerCase();
            const filtered = notes.filter((note) =>
                note.title.toLowerCase().includes(lowerCaseQuery)
            );
            setFilteredNotes(filtered);
        }
    };

    const handleSaveNote = async (note) => {
        try {
            if (note.id) {
                const response = await axios.put(API_BASE_URL, note);
                setNotes((prevNotes) =>
                    prevNotes.map((n) => (n.id === note.id ? response.data : n))
                );
            } else {
                const response = await axios.post(API_BASE_URL, note);
                setNotes((prevNotes) => [...prevNotes, response.data]);
            }
            setFilteredNotes(notes);
        } catch (error) {
            console.error("Ошибка при сохранении заметки:", error);
        }
    };

    const handleDeleteNote = async (id) => {
        try {
            await axios.delete(`${API_BASE_URL}/${id}`);
            setNotes((prevNotes) => prevNotes.filter((note) => note.id !== id));
            setFilteredNotes((prevNotes) => prevNotes.filter((note) => note.id !== id));
        } catch (error) {
            console.error("Ошибка при удалении заметки:", error);
        }
    };

    const handleCreateNote = () => {
        setCurrentNote({ title: "", content: "" });
        setIsModalOpen(true);
    };

    const handleEditNote = (note) => {
        setCurrentNote(note);
        setIsModalOpen(true);
    };

    return (
        <div className="app">
            <h1>Мои заметки</h1>
            <div className="actions">
                <button className="btn-primary mb-2" onClick={handleCreateNote}>
                    Создать заметку
                </button>
                <SearchBar notes={notes} onSearch={handleSearch} />
            </div>
            <NoteList
                notes={filteredNotes}
                onEdit={handleEditNote}
                onDelete={handleDeleteNote}
            />
            {isModalOpen && (
                <NoteModal
                    note={currentNote}
                    onClose={() => setIsModalOpen(false)}
                    onSave={handleSaveNote}
                />
            )}
        </div>
    );
};

export default App;