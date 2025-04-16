import React, { createContext, useContext, useState, useEffect } from "react";
import axios from "axios";

const AuthContext = createContext();

export function useAuth() {
    return useContext(AuthContext);
}

export function AuthProvider({ children }) {
    const [currentUser, setCurrentUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

    const apiUrl = process.env.REACT_APP_API_URL || "http://localhost:8080";

    useEffect(() => {
        // ローカルストレージからユーザー情報を取得
        const user = localStorage.getItem("user");
        const token = localStorage.getItem("token");

        if (user && token) {
            setCurrentUser(JSON.parse(user));
            // トークン付きのAxiosインスタンスを設定
            axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        }

        setLoading(false);
    }, []);

    // 登録機能
    const register = async (username, email, password) => {
        try {
            setError("");
            const response = await axios.post(
                `${apiUrl}/api/v1/auth/register`,
                {
                    username,
                    email,
                    password,
                }
            );

            // ユーザー情報とトークンを保存
            setCurrentUser({
                username: response.data.username,
                email: response.data.email,
            });

            localStorage.setItem("token", response.data.token);
            localStorage.setItem(
                "user",
                JSON.stringify({
                    username: response.data.username,
                    email: response.data.email,
                })
            );

            // トークン付きのAxiosインスタンスを設定
            axios.defaults.headers.common[
                "Authorization"
            ] = `Bearer ${response.data.token}`;

            return response.data;
        } catch (err) {
            setError(err.response?.data?.error || "Registration failed");
            throw err;
        }
    };

    // ログイン機能
    const login = async (email, password) => {
        try {
            setError("");
            const response = await axios.post(`${apiUrl}/api/v1/auth/login`, {
                email,
                password,
            });

            // ユーザー情報とトークンを保存
            setCurrentUser({
                username: response.data.username,
                email: response.data.email,
            });

            localStorage.setItem("token", response.data.token);
            localStorage.setItem(
                "user",
                JSON.stringify({
                    username: response.data.username,
                    email: response.data.email,
                })
            );

            // トークン付きのAxiosインスタンスを設定
            axios.defaults.headers.common[
                "Authorization"
            ] = `Bearer ${response.data.token}`;

            return response.data;
        } catch (err) {
            setError(err.response?.data?.error || "Login failed");
            throw err;
        }
    };

    // ログアウト機能
    const logout = () => {
        setCurrentUser(null);
        localStorage.removeItem("token");
        localStorage.removeItem("user");
        delete axios.defaults.headers.common["Authorization"];
    };

    const value = {
        currentUser,
        register,
        login,
        logout,
        error,
    };

    return (
        <AuthContext.Provider value={value}>
            {!loading && children}
        </AuthContext.Provider>
    );
}
