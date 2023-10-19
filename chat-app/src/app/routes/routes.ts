import { Routes } from '@angular/router';
import { AuthGuard } from '../core/guards/auth.guard';
import { APP_ROLES } from '../core/constants/roles';

export const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('../modules/home/home.module').then((m) => m.HomeModule),
    canActivate: [AuthGuard],
    data: {
      roles: [APP_ROLES.ADMIN, APP_ROLES.USER],
    },
  },
  {
    path: 'login',
    loadChildren: () => import('../modules/login/login.module').then((m) => m.LoginModule),
  },
  {
    path: 'error',
    loadChildren: () => import('../modules/error-pages/error-pages.module').then((m) => m.ErrorPagesModule),
    data: {
      breadcrumbHide: true,
    },
  },
  {
    path: '**',
    redirectTo: 'error/not-found',
  },
];
