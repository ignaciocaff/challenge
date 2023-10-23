import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../models';
import { Observable } from 'rxjs';

/**
 * Divide the http logic from the business logic, of the session service.
 * Use this approach to make the code more readable and maintainable, whenever it's possible.
 */
@Injectable({
  providedIn: 'root',
})
export class SessionHttpService {
  constructor(private readonly http: HttpClient) {}

  authUser(userName: string, password: string) {
    return this.http.post<User>('/api/auth', { userName, password });
  }

  signup(user: User) {
    return this.http.post<User>('/api/auth/signup', user);
  }

  me() {
    return this.http.get('/api/auth/me');
  }

  logout(): Observable<any> {
    return this.http.get('/api/auth/logout');
  }
}
