import React, { useState, useContext } from "react";
import {
  Container,
  TextField,
  Button,
  Typography,
  Box,
  Alert,
} from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import { login, register } from "../api/api";
import { AuthContext } from "../App";

const LoginRegister: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const isRegister = location.pathname === "/register";
  const { setToken } = useContext(AuthContext);

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const apiCall = isRegister ? register : login;
      const response = await apiCall({ username, password });
      if (!response.ok) {
        const errorMessage = await response.text();
        setError(errorMessage || "Error occurred");
        return;
      }
      // Get JWT token from header (the API returns it in the Authorization header)
      const respBody = await response.json();
      const token = response.headers.get("Authorization");
      console.log(respBody);
      if (token) {
        // Remove "Bearer " prefix if present
        const tokenValue = token.startsWith("Bearer ")
          ? token.substring(7)
          : token;
        setToken(tokenValue);
        navigate("/");
      } else {
        setError("Token not received");
      }
    } catch (err: unknown) {
      setError((err as Error).message || "Error occurred");
    }
  };

  return (
    <Container maxWidth="sm">
      <Box mt={8} display="flex" flexDirection="column" alignItems="center">
        <Typography component="h1" variant="h5">
          {isRegister ? "Register" : "Login"}
        </Typography>
        {error && (
          <Alert severity="error" sx={{ mt: 2, width: "100%" }}>
            {error}
          </Alert>
        )}
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            label="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            label="Password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button type="submit" fullWidth variant="contained" sx={{ mt: 3 }}>
            {isRegister ? "Register" : "Login"}
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default LoginRegister;
