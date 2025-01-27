import React, { useState } from "react";
import PostList from "./components/PostList";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import PostForm from "./components/PostForm";

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(Boolean(localStorage.getItem("jewelry123")));

  return (
    <div>
      <h1>Welcome to the Admin Panel</h1>
      {!isAuthenticated ? (
        <>
          <RegisterForm />
          <LoginForm setIsAuthenticated={setIsAuthenticated} />
        </>
      ) : (
        <>
          <PostForm isAuthenticated={isAuthenticated} />
          <PostList />
        </>
      )}
    </div>
  );
};

export default App;

