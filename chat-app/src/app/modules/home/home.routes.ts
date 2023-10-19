import { Routes } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';

export const routes: Routes = [
  {
    path: '',
    component: DashboardComponent,
    data: {
      breadcrumbHide: true,
    },
  },
  {
    path: 'rooms/:id',
    loadChildren: () => import('../room/room.module').then((m) => m.RoomModule),
  },
];
