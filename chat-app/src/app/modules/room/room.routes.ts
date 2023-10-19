import { Routes } from '@angular/router';
import { RoomComponent } from './room.component';
import { AuthGuard } from 'src/app/core/guards/auth.guard';
import { APP_ROLES } from 'src/app/core/constants/roles';

export const routes: Routes = [
  {
    path: '',
    component: RoomComponent,
    data: {
      pageTitle: 'Room',
      breadcrumb: 'Room',
      roles: [APP_ROLES.ADMIN, APP_ROLES.USER],
    },
    canActivate: [AuthGuard],
  },
];
