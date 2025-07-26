import axios from "axios";

export const API_BASE = import.meta.env.VITE_API_BASE ?? "http://localhost:8080";

const api = axios.create({
  baseURL: API_BASE,
  withCredentials: true,
});

// Add Authorization header from token in localStorage
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// For navigation callback
let navigate: ((path: string) => void) | null = null;
export const setNavigate = (nav: (path: string) => void) => {
  navigate = nav;
};

// Response interceptor for global 401 handling
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const status = error.response?.status
    if (status === 401) {
      console.warn("Global 401 Unauthorized detected");

      // Clear auth info if you want
      localStorage.removeItem("token");

      if (navigate) {
        navigate("/login");
      } else {
        // fallback - reload page or hard redirect
        window.location.href = "/login";
      }
    } else if (status == 403) {
      console.warn("Global 403 Forbidden detected");

      if (navigate) {
        navigate("/unauthorized");
      } else {
        window.location.href = "/unauthorized";
      } console.warn("Global 403 Forbidden detected");

      if (navigate) {
        navigate("/unauthorized");
      } else {
        window.location.href = "/unauthorized";
      }
    }
    return Promise.reject(error);
  }
);

export default api;
