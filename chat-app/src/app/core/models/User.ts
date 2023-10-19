import { Role } from './Role';

export interface User {
  id?: number;
  email?: string;
  username?: string;
  password?: string;
  roles: Role[];
}
