import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";
import { buildCredentials } from "../models/Credentials";


export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    const credentials = buildCredentials({ username, password })
    try {
      const res = await api.post(`api/login`, credentials);

      const token = res.data.token;
      localStorage.setItem("token", token);

      localStorage.setItem("user", JSON.stringify(res.data.user));

      navigate("/home"); // Redirect to homepage
    } catch (err) {
      setError("Invalid username or password.");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-6 text-center">Log In</h2>
        {error && <div className="text-red-500 mb-4">{error}</div>}
        <form onSubmit={handleLogin} className="space-y-4">
          <div>
            <label className="block mb-1 text-sm font-medium text-gray-700">
              Username
            </label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-green-400"
              required
            />
          </div>
          <div>
            <label className="block mb-1 text-sm font-medium text-gray-700">
              Password
            </label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full border px-3 py-2 rounded focus:outline-none focus:ring-2 focus:ring-green-400"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-green-500 text-white py-2 rounded hover:bg-green-600 transition"
          >
            Log In
          </button>
        </form>
        <p className="mt-4 text-sm text-center">
          Don’t have an account?{" "}
          <a href="/register" className="text-green-500 hover:underline">
            Register
          </a>
        </p>
      </div>
    </div>
  );
}
