// src/pages/Stories.js
import React, { useState, useEffect, useContext } from 'react';
import API from '../services/api';
import { AuthContext } from '../contexts/AuthContext';

const Stories = () => {
  const { auth } = useContext(AuthContext);
  const [stories, setStories] = useState([]);
  const [newStory, setNewStory] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const fetchStories = () => {
    API.get('/stories')
      .then((res) => setStories(res.data.data))
      .catch((err) => console.error(err));
  };

  useEffect(() => {
    fetchStories();
    // آپشنال: تازه‌سازی داستان‌ها هر ۴۸ ساعت
    const interval = setInterval(fetchStories, 1000 * 60 * 60 * 48);
    return () => clearInterval(interval);
  }, []);

  const handleChange = (e) => {
    setNewStory(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!auth.user) {
      setError('You must be logged in to post a story.');
      return;
    }

    const payload = {
      content: newStory,
    };

    API.post('/stories', payload)
      .then((res) => {
        setSuccess('Story posted successfully!');
        setError('');
        setNewStory('');
        fetchStories();
      })
      .catch((err) => {
        setError(err.response?.data?.error || 'Failed to post story');
        setSuccess('');
      });
  };

  return (
    <div style={styles.container}>
      <h1>Stories</h1>
      {auth.user && (
        <form onSubmit={handleSubmit} style={styles.form}>
          <textarea
            name="story"
            placeholder="Share your purchase story..."
            value={newStory}
            onChange={handleChange}
            required
            style={styles.textarea}
          />
          {error && <p style={styles.error}>{error}</p>}
          {success && <p style={styles.success}>{success}</p>}
          <button type="submit" style={styles.button}>
            Post Story
          </button>
        </form>
      )}
      <div style={styles.stories}>
        {stories.length === 0 ? (
          <p>No stories available.</p>
        ) : (
          stories.map((story) => (
            <div key={story.ID} style={styles.story}>
              <p>{story.Content}</p>
              <p style={styles.author}>- {story.User.Username}</p>
              <p style={styles.date}>
                {new Date(story.CreatedAt).toLocaleString()}
              </p>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

const styles = {
  container: {
    padding: '20px',
    maxWidth: '800px',
    margin: '0 auto',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    marginBottom: '20px',
  },
  textarea: {
    padding: '10px',
    marginBottom: '10px',
    fontSize: '16px',
    height: '100px',
  },
  button: {
    padding: '10px',
    backgroundColor: '#4CAF50',
    color: '#fff',
    border: 'none',
    cursor: 'pointer',
    fontSize: '16px',
  },
  stories: {
    display: 'flex',
    flexDirection: 'column',
    gap: '15px',
  },
  story: {
    padding: '15px',
    border: '1px solid #ccc',
    borderRadius: '5px',
  },
  author: {
    fontStyle: 'italic',
    color: '#555',
  },
  date: {
    fontSize: '12px',
    color: '#999',
  },
  error: {
    color: 'red',
    marginBottom: '10px',
  },
  success: {
    color: 'green',
    marginBottom: '10px',
  },
};

export default Stories;

