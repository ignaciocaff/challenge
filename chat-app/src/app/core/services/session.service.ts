import { Injectable } from '@angular/core';
import { Observable, BehaviorSubject, of } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { SessionHttpService } from './session-http.service';
import { Role, User } from '../models';
import { environment } from 'src/environments/environment';

/**
 * This service is used to handle the application session.
 */
@Injectable({
  providedIn: 'root',
})
export class SessionService {
  private loggedUserSubject: BehaviorSubject<User | undefined> = new BehaviorSubject<User | undefined>(undefined);

  constructor(private readonly sessionHttp: SessionHttpService) { }

  me(): Observable<any> {
    return this.sessionHttp.me();
  }

  logout(): Observable<void> {
    return this.sessionHttp.logout().pipe(
      map(() => {
        localStorage.clear();
        this.loggedUserSubject.next(undefined);
      })
    );
  }
  setLoggedUser(user: User): void {
    this.loggedUserSubject.next(user);
  }

  getUserRoles(): Observable<Role[]> {
    return this.loggedUserSubject.pipe(map((user) => user?.roles ?? []));
  }

  getLoggedUser(): Observable<User | undefined> {
    return this.loggedUserSubject.asObservable();
  }
}
