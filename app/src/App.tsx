import React, { useState, createContext } from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import LoginRegister from "./pages/LoginRegister";
import Home from "./pages/Home";
import Reservation from "./pages/Reservation";
import MyReservations from "./pages/MyReservations";
import Navbar from "./components/Navbar";

export interface AuthContextType {
  token: string | null;
  setToken: (token: string | null) => void;
}

export const AuthContext = createContext<AuthContextType>({
  token: null,
  setToken: () => {},
});

const App: React.FC = () => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token"),
  );

  const handleSetToken = (token: string | null) => {
    if (token) {
      localStorage.setItem("token", token);
    } else {
      localStorage.removeItem("token");
    }
    setToken(token);
  };

  return (
    <AuthContext.Provider value={{ token, setToken: handleSetToken }}>
      <BrowserRouter>
        <Navbar />
        <Routes>
          <Route path="/login" element={<LoginRegister />} />
          <Route path="/register" element={<LoginRegister />} />
          {/* Protected Routes */}
          <Route
            path="/"
            element={token ? <Home /> : <Navigate to="/login" replace />}
          />
          <Route
            path="/reservation/:projectionId"
            element={token ? <Reservation /> : <Navigate to="/login" replace />}
          />
          <Route
            path="/reservations"
            element={
              token ? <MyReservations /> : <Navigate to="/login" replace />
            }
          />
        </Routes>
      </BrowserRouter>
    </AuthContext.Provider>
  );
};

export default App;
