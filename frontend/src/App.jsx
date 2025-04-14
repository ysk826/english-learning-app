import React, { useState, useEffect } from "react";
import "./App.css";

function App() {
    const [status, setStatus] = useState("Loading...");

    useEffect(() => {
        // Check backend health
        fetch("http://localhost:8080/health")
            .then((response) => response.json())
            .then((data) => {
                setStatus(
                    data.status === "ok"
                        ? "Backend is connected!"
                        : "Backend connection issue"
                );
            })
            .catch((error) => {
                console.error("Error connecting to backend:", error);
                setStatus("Failed to connect to backend");
            });
    }, []);

    return (
        <div className="App">
            <header className="App-header">
                <h1>English Learning App</h1>
                <p>Backend Status: {status}</p>
                <div>
                    <h2>Features coming soon:</h2>
                    <ul>
                        <li>User Registration</li>
                        <li>Word Collection</li>
                        <li>Sentence Creation</li>
                    </ul>
                </div>
            </header>
        </div>
    );
}

export default App;
