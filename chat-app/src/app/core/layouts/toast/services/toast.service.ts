import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';
import { ToastData } from '../models/ToastData';

@Injectable({
  providedIn: 'root',
})
export class ToastService {
  private alertSubject = new Subject<ToastData>();

  constructor() {}

  getAlert(): Observable<ToastData> {
    return this.alertSubject.asObservable();
  }

  show(alert: ToastData): void {
    this.alertSubject.next(alert);
  }

  showError(title: string, message: string): void {
    this.alertSubject.next({
      type: 'error',
      title,
      message,
    });
  }

  showSuccess(title: string, message: string): void {
    this.alertSubject.next({
      type: 'success',
      title,
      message,
    });
  }
}
