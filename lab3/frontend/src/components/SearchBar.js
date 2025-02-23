import React, { useState } from "react";
import Autosuggest from "react-autosuggest";
import "./styles/SearchBar.css";

// Функция для фильтрации заметок по введенному тексту
const getSuggestions = (value, notes) => {
    const inputValue = value.trim().toLowerCase();
    const inputLength = inputValue.length;

    return inputLength < 2
        ? [] // Показываем предложения только после ввода 2+ символов
        : notes.filter(
            (note) => note.title.toLowerCase().includes(inputValue)
        );
};

// Функция для получения текста из объекта предложения
const getSuggestionValue = (suggestion) => suggestion.title;

const SearchBar = ({ notes, onSearch }) => {
    const [value, setValue] = useState("");
    const [suggestions, setSuggestions] = useState([]);

    const onChange = (event, { newValue }) => {
        setValue(newValue);
        onSearch(newValue); // Передаем введенное значение в родительский компонент
    };

    const onSuggestionsFetchRequested = ({ value }) => {
        setSuggestions(getSuggestions(value, notes));
    };

    const onSuggestionsClearRequested = () => {
        setSuggestions([]);
    };

    const renderSuggestion = (suggestion) => <div>{suggestion.title}</div>;

    return (
        <Autosuggest
            suggestions={suggestions}
            onSuggestionsFetchRequested={onSuggestionsFetchRequested}
            onSuggestionsClearRequested={onSuggestionsClearRequested}
            getSuggestionValue={getSuggestionValue}
            renderSuggestion={renderSuggestion}
            inputProps={{
                placeholder: "Поиск заметок...",
                value,
                onChange,
            }}
        />
    );
};

export default SearchBar;