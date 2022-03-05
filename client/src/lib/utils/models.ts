export interface SurfSession {
  id: number;
  description: string;
  datetime: string;
  swell: string;
  wind: string;
  tide: string;
  surf_spot: string;
  spot_location: string;
  spot_region: string;
  creator_id: number;
  surfboard_id?: number;
  is_public: boolean;
}

export interface Favorite {
  id: number;
  surf_session_id: number;
  user_id: number;
}

export interface Comment {
  id: number;
  surf_session_id: number;
  content: string;
  commenter_id: number;
}

export interface SavedSpot {
  id: number;
  spot_name: string;
  spot_region: string;
  user_id: number;
}

export interface User {
  id: number;
  username: string;
  first_name: string;
  last_name: string;
  email: string;
}
