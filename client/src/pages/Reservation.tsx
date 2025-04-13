import React, { useState, useContext } from "react";
import {
  Container,
  TextField,
  Button,
  Typography,
  Box,
  Alert,
} from "@mui/material";
import { useParams, useNavigate } from "react-router-dom";
import { createReservation } from "../api/api";
import { AuthContext } from "../App";

const Reservation: React.FC = () => {
  const { projectionId } = useParams<{ projectionId: string }>();
  const { token } = useContext(AuthContext);
  const navigate = useNavigate();
  const [seats, setSeats] = useState<number>(1);
  const [error, setError] = useState<string | null>(null);

  const handleReserve = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!token) return;
    try {
      await createReservation(token, {
        projection_id: parseInt(projectionId!),
        seats,
      });
      navigate("/reservations");
    } catch (err: unknown) {
      setError((err as Error).message || "Reservation failed");
    }
  };

  return (
    <Container maxWidth="sm">
      <Box mt={8}>
        <Typography variant="h5" gutterBottom>
          Reserve Seats for Projection {projectionId}
        </Typography>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}
        <Box component="form" onSubmit={handleReserve}>
          <TextField
            label="Number of Seats"
            type="number"
            value={seats}
            onChange={(e) => setSeats(parseInt(e.target.value))}
            fullWidth
            required
            sx={{ mb: 2 }}
          />
          <Button type="submit" variant="contained" fullWidth>
            Reserve
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Reservation;
