// frontend/src/components/StoryList.js
import React, { useEffect, useState } from 'react';
import api from '../services/api';
import StoryCard from './StoryCard';

const StoryList = () => {
    const [stories, setStories] = useState([]);

    const fetchStories = async () => {
        try {
            const response = await api.get('/stories');
            setStories(response.data.data);
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        fetchStories();
    }, []);

    return (
        <div>
            <h2>Purchase Stories</h2>
            {stories.map(story => (
                <StoryCard key={story.id} story={story} />
            ))}
        </div>
    );
};

export default StoryList;

