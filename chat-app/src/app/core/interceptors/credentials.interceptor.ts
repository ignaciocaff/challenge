import { Injectable } from '@angular/core';
import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class CredentialsInterceptor implements HttpInterceptor {
  constructor() {}
  intercept(req: HttpRequest<undefined>, next: HttpHandler): Observable<HttpEvent<undefined>> {
    return next.handle(req);
  }
}
