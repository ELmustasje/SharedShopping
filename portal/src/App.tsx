import { BrowserRouter, Routes, Route, useNavigate } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ShoppingListPage from "./pages/ShoppingListPage";
import RegisterPage from "./pages/RegisterPage";
import LoginPage from "./pages/LoginPage";
import { useEffect, useState } from "react";
import { setNavigate } from "./api/axios";
import UnauthorizedPage from "./pages/ERROR/UnauthorizedPage";
import LoginRequiredPage from "./pages/ERROR/LoginRequiredPage";
import NotFoundPage from "./pages/ERROR/NotFoundPage";
import { UserContext } from "./contexts/userContext.tsx";
import { buildUser, type User } from "./models/User"

function NavigatorSetter() {
  const navigate = useNavigate();

  useEffect(() => {
    setNavigate(navigate);
  }, [navigate]);

  return null;
}

export default function App() {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    // fetch from localStorage or API
    const userData = localStorage.getItem("user");
    if (userData) setUser(JSON.parse(userData));
  }, []);

  return (
    <UserContext.Provider value={user}>
      <BrowserRouter>
        <NavigatorSetter />
        <Routes>
          <Route path="/home" element={<HomePage />} />
          <Route path="/list/:id" element={<ShoppingListPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/unauthorized" element={<UnauthorizedPage />} />
          <Route path="/login-required" element={<LoginRequiredPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </BrowserRouter>
    </UserContext.Provider>
  );
}
