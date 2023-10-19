import { Component, OnInit } from '@angular/core';
import { SessionService } from '../../services/session.service';
import { User } from '../../models';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
})
export class NavbarComponent implements OnInit {
  user: User | undefined;
  constructor(private readonly sessionService: SessionService,
    private readonly router: Router
  ) { }

  ngOnInit(): void {
    this.sessionService.getLoggedUser().subscribe((user) => {
      this.user = user;
    });
  }

  public logout(): void {
    this.sessionService.logout().subscribe({
      next: () => {
        this.router.navigate(['/login']);
      },
      error: () => {
        console.error('Error during logout:');
      },
    });
  }

  public getLoggedUser(): Observable<User | undefined> {
    return this.sessionService.getLoggedUser();
  }
}
