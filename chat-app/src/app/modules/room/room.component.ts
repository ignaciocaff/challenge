import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { WebSocketService } from './services/websocket.service';
import { SessionService } from 'src/app/core/services/session.service';
import { User } from 'src/app/core/models';
import { RoomService } from './services/room.service';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
})
export class RoomComponent implements OnInit {
  @ViewChild('messageContainer') messageContainer: ElementRef | undefined;

  roomId?: string;
  newMessage: string = '';
  user: User | undefined;

  messages: any = [];

  constructor(private route: ActivatedRoute,
    private webSocketService: WebSocketService,
    private readonly sessionService: SessionService,
    private readonly roomService: RoomService) { }

  ngAfterViewChecked() {
    this.scrollToBottom();
  }
  scrollToBottom() {
    if (this.messageContainer) this.messageContainer.nativeElement.scrollTop = this.messageContainer.nativeElement.scrollHeight;
  }

  ngOnInit(): void {
    this.route.params.subscribe((params: { [x: string]: string; }) => {
      this.roomId = params['id'];
      this.roomService.messages(this.roomId).subscribe({
        next: (storedMessages: any) => {
          this.messages = storedMessages ?? [];
          this.webSocketService.connect(this.roomId!).subscribe({
            next: (message: any) => {
              this.messages.push(message);
            },
            error: (error: any) => {
              console.error(error);
            },
          });
        },
        error: (error: any) => {
          console.error(error);
        },
      });
    });
    this.sessionService.getLoggedUser().subscribe((user: User | undefined) => {
      this.user = user;
    });
  }
  sendMessage() {
    if (this.newMessage.trim() !== '') {
      const payload = {
        sender: {
          username: this.user?.username,
          id: this.user?.id
        },
        date: new Date(),
        text: this.newMessage
      };
      this.webSocketService.sendMessage(payload);
      this.newMessage = '';
    }
  }

  getCurrentTime(): string {
    const date = new Date();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    const seconds = date.getSeconds();
    return `${hours}:${minutes < 10 ? '0' : ''}${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
  }
}
