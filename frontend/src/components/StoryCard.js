// frontend/src/components/StoryCard.js
import React from 'react';

const StoryCard = ({ story }) => {
    return (
        <div className="story-card">
            <h4>{story.user.username}</h4>
            <p>{story.content}</p>
            <small>{new Date(story.created_at).toLocaleString()}</small>
        </div>
    );
};

export default StoryCard;

