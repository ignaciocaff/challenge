import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WebSocketService {
  private socket: WebSocket | undefined;

  constructor() {
  }

  public connect(roomId: string): Observable<any> {
    this.socket = new WebSocket(`ws://localhost:3000/ws/${roomId}`,);

    return new Observable(observer => {
      this.socket!.onmessage = (event) => {
        const data = JSON.parse(event.data);
        observer.next(data);
      };
      this.socket!.onerror = (event) => observer.error(event);
      this.socket!.onclose = () => observer.complete();
    });
  }

  public sendMessage(payload: any): void {
    const message = JSON.stringify(payload);
    this.socket!.send(message);
  }
}
