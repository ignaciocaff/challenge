import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { APP_ROLES } from 'src/app/core/constants/roles';
import { SessionService } from 'src/app/core/services/session.service';
import { DashboardService } from '../service/dashboard.service';
import { Room } from 'src/app/core/models';

@Component({
  templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {
  readonly APP_ROLES = APP_ROLES;
  rooms: Room[] = [];
  room: Room = {
    name: '',
  };
  showModal = false;
  isDialogOpen: boolean = false;
  constructor(
    private readonly sessionService: SessionService,
    private readonly dashboardService: DashboardService,
    private router: Router
  ) {
    this.sessionService.getLoggedUser().subscribe((user: any) => {
      if (!user) return;
    });
  }

  ngOnInit(): void {
    this.dashboardService.rooms().subscribe((rooms: Room[]) => {
      this.rooms = rooms;
    });
  }

  navigateToRoom(roomId: number) {
    this.router.navigate(['/rooms', roomId]);
  }

  openDialog() {
    this.isDialogOpen = true;
  }

  closeDialog() {
    this.isDialogOpen = false;
  }
  goBack(): void {
    this.router.navigate(['/']);
  }

  toggleModal(): void {
    this.showModal = !this.showModal;
  }

  create() {
    this.dashboardService.create(this.room).subscribe({
      next: () => {
        alert('Room created successfully');
        this.dashboardService.rooms().subscribe((rooms: Room[]) => {
          this.rooms = rooms;
        });
      },
      error: (err) => {
        alert(err?.error?.message);
        return;
      },
    });
  }
}
