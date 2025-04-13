import React, { useEffect, useState, useContext } from "react";
import {
  Container,
  Grid,
  Card,
  CardContent,
  Typography,
  Button,
  CardActions,
} from "@mui/material";
import { getProjections, Projection } from "../api/api";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../App";

const Home: React.FC = () => {
  const [projections, setProjections] = useState<Projection[]>([]);
  const { token } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (token) {
      getProjections(token)
        .then((data) => setProjections(data || []))
        .catch((err) => console.error(err));
    }
  }, [token]);

  return (
    <Container sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        Movie Projections
      </Typography>
      <Grid container spacing={2}>
        {projections.map((projection) => (
          <Grid item xs={12} sm={6} md={4} key={projection.id}>
            <Card>
              <CardContent>
                <Typography variant="h6">
                  Movie ID: {projection.movie_id}
                </Typography>
                <Typography color="textSecondary">
                  Cinema: {projection.cinema}
                </Typography>
                <Typography color="textSecondary">
                  Showtime: {new Date(projection.showtime).toLocaleString()}
                </Typography>
              </CardContent>
              <CardActions>
                <Button
                  size="small"
                  onClick={() => navigate(`/reservation/${projection.id}`)}
                >
                  Reserve
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default Home;
