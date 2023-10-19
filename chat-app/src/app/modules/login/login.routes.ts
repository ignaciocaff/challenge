import { Routes } from '@angular/router';
import { LoginComponent } from './login.component';
import { AuthGuard } from 'src/app/core/guards/auth.guard';
import { APP_ROLES } from 'src/app/core/constants/roles';

export const routes: Routes = [
  {
    path: '',
    component: LoginComponent,
    data: {
      pageTitle: 'Login',
      showBreadcrumb: false,
      roles: [],
    },
    canActivate: [],
  },
];
