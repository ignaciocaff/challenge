import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/app/core/models';
import { SessionHttpService } from 'src/app/core/services/session-http.service';
import { SessionService } from 'src/app/core/services/session.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(
    private authService: SessionHttpService,
    private readonly sessionService: SessionService,
    private readonly router: Router
  ) {}

  login(): void {
    this.authService.authUser(this.username, this.password).subscribe({
      next: (user: User) => {
        this.sessionService.setLoggedUser(user);
        this.router.navigate(['/']);
      },
      error: () => {
        alert('User or password incorrect');
        return;
      },
    });
  }
}
