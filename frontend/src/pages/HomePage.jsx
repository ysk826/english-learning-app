import React, { useState, useEffect } from "react";
import { useAuth } from "../contexts/AuthContext";

function HomePage() {
    const { currentUser } = useAuth();
    const [status, setStatus] = useState("Loading...");
    const apiUrl = process.env.REACT_APP_API_URL || "http://localhost:8080";

    useEffect(() => {
        // Check backend health
        fetch(`${apiUrl}/health`)
            .then((response) => response.json())
            .then((data) => {
                setStatus(
                    data.status === "ok"
                        ? `Backend is connected! ${data.message || ""}`
                        : `Backend connection issue: ${data.message || ""}`
                );
            })
            .catch((error) => {
                console.error("Error connecting to backend:", error);
                setStatus("Failed to connect to backend");
            });
    }, [apiUrl]);

    return (
        <div className="main-content home-container">
            <h1 className="page-title">Welcome to English Learning App</h1>
            {currentUser ? (
                <div>
                    <p>Hello, {currentUser.username}!</p>
                    <p>You are logged in with {currentUser.email}</p>
                </div>
            ) : (
                <p>Please register or login to start learning English.</p>
            )}

            <p>Backend Status: {status}</p>

            <h2>Features</h2>
            <ul className="feature-list">
                <li className="feature-item">
                    <h3>Word Collection</h3>
                    <p>
                        Save interesting English words that you want to
                        remember.
                    </p>
                </li>
                <li className="feature-item">
                    <h3>Sentence Creation</h3>
                    <p>Create sentences using the words you've collected.</p>
                </li>
            </ul>
        </div>
    );
}

export default HomePage;
