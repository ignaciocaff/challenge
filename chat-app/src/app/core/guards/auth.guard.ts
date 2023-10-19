import { Injectable, isDevMode } from '@angular/core';
import { ActivatedRouteSnapshot, Router } from '@angular/router';
import { map } from 'rxjs';
import { SessionService } from '../services/session.service';
import { User } from '../models';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard {
  constructor(
    private readonly sessionService: SessionService,
    private readonly router: Router
  ) { }

  canActivate() {
    return this.sessionService.getLoggedUser().pipe(
      map((user) => {
        if (!user) {
          return this.sessionService.me().subscribe({
            next: (user: User) => {
              this.sessionService.setLoggedUser(user);
              return true
            },
            error: () => {
              this.router.navigate(['/login']);
              return false;
            },
          });
        }
        return true
      })
    );
  }
}
