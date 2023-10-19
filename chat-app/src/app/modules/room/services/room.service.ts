import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Message, Room } from 'src/app/core/models';

@Injectable({
  providedIn: 'root',
})
export class RoomService {
  constructor(private readonly http: HttpClient) { }

  messages(roomId: string): Observable<Message[]> {
    return this.http.get<Message[]>(`/api/rooms/${roomId}/messages`);
  }
}
