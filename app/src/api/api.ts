export const API_BASE_URL = "http://localhost:8080/api";

export interface UserRegistration {
  username: string;
  password: string;
}

export interface UserLogin {
  username: string;
  password: string;
}

export interface Movie {
  id: number;
  title: string;
  description?: string;
  duration: number;
}

export interface Projection {
  id: number;
  movie_id: number;
  cinema: string;
  showtime: string;
}

export interface Reservation {
  id: number;
  user_id: number;
  projection_id: number;
  seats: number;
}

export interface ReservationInput {
  projection_id: number;
  seats: number;
}

export async function login(data: UserLogin): Promise<Response> {
  console.log(data);
  return fetch(`${API_BASE_URL}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

export async function register(data: UserRegistration): Promise<Response> {
  return fetch(`${API_BASE_URL}/auth/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

export async function getProjections(token: string): Promise<Projection[]> {
  const response = await fetch(`${API_BASE_URL}/projections`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch projections");
  }
  return response.json();
}

export async function getReservations(token: string): Promise<Reservation[]> {
  const response = await fetch(`${API_BASE_URL}/reservations`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch reservations");
  }
  return response.json();
}

export async function createReservation(
  token: string,
  data: ReservationInput,
): Promise<Reservation> {
  const response = await fetch(`${API_BASE_URL}/reservations`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    throw new Error("Failed to create reservation");
  }
  return response.json();
}
