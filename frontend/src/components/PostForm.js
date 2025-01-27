import React, { useState, useEffect } from "react";
import axios from "axios";

const PostForm = ({ isAuthenticated }) => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [token, setToken] = useState(localStorage.getItem("jewelry123"));
  const [error, setError] = useState(null);
  const [successMessage, setSuccessMessage] = useState("");
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const checkAdminRole = async () => {
      try {
        const response = await axios.get("http://localhost:8080/api/check-admin", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setIsAdmin(response.data.isAdmin);
      } catch (err) {
        setIsAdmin(false);
      }
    };
    if (token) {
      checkAdminRole();
    }
  }, [token]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post(
        "http://localhost:8080/api/posts/",
        { title, content },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setSuccessMessage("Post created successfully!");
      setTitle("");
      setContent("");
      setError(null);
    } catch (err) {
      setError("Error creating post. Ensure you have the admin role.");
      setSuccessMessage("");
    }
  };

  return (
    <div>
      <h2>Create New Post</h2>
      {successMessage && <p style={{ color: "green" }}>{successMessage}</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
      {isAuthenticated && isAdmin && (
        <form onSubmit={handleSubmit}>
          <div>
            <label>Title:</label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          <div>
            <label>Content:</label>
            <textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
              required
            />
          </div>
          <button type="submit">Create Post</button>
        </form>
      )}
    </div>
  );
};

export default PostForm;

