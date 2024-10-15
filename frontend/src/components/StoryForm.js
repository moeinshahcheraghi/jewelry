// frontend/src/components/StoryForm.js
import React, { useState } from 'react';
import api from '../services/api';
import { useHistory } from 'react-router-dom';

const StoryForm = ({ onStoryCreated }) => {
    const [content, setContent] = useState('');
    const history = useHistory();
    const token = localStorage.getItem('token');

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!token) {
            alert('Please login to post a story');
            history.push('/login');
            return;
        }
        try {
            const response = await api.post('/stories', { content });
            alert('Story added successfully');
            setContent('');
            onStoryCreated(response.data.data);
        } catch (error) {
            alert(error.response.data.error || 'Error adding story');
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <textarea
                value={content}
                onChange={(e) => setContent(e.target.value)}
                placeholder="Share your purchase story..."
                required
            ></textarea>
            <button type="submit">Post Story</button>
        </form>
    );
};

export default StoryForm;

