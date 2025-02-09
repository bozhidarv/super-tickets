import React, { useEffect, useState, useContext } from "react";
import {
  Container,
  Typography,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";
import { getReservations, Reservation } from "../api/api";
import { AuthContext } from "../App";

const MyReservations: React.FC = () => {
  const { token } = useContext(AuthContext);
  const [reservations, setReservations] = useState<Reservation[]>([]);

  useEffect(() => {
    if (token) {
      getReservations(token)
        .then((data) => setReservations(data))
        .catch((err) => console.error(err));
    }
  }, [token]);

  return (
    <Container sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        My Reservations
      </Typography>
      <List>
        {reservations.map((reservation) => (
          <ListItem key={reservation.id}>
            <ListItemText
              primary={`Reservation ID: ${reservation.id}`}
              secondary={`Projection ID:
          ${reservation.projection_id} — Seats: ${reservation.seats}`}
            />
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default MyReservations;
