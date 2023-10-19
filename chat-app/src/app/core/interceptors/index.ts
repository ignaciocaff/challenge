import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { LoadingInterceptor } from './ loading.interceptor';
import { CredentialsInterceptor } from './credentials.interceptor';

export const httpInterceptorProviders = [
  {
    provide: HTTP_INTERCEPTORS,
    useClass: LoadingInterceptor,
    multi: true,
  },
  {
    provide: HTTP_INTERCEPTORS,
    useClass: CredentialsInterceptor,
    multi: true,
  },
];
