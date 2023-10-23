import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Room } from 'src/app/core/models';

@Injectable({
  providedIn: 'root',
})
export class DashboardService {
  constructor(private readonly http: HttpClient) {}

  rooms() {
    return this.http.get<Room[]>('/api/rooms');
  }

  create(room: Room) {
    return this.http.post<Room[]>('/api/rooms/create', room);
  }
}
