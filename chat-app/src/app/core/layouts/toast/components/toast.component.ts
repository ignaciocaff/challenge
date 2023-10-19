import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-toast',
  templateUrl: './toast.component.html',
})
export class ToastComponent {
  @Input() title?: string;
  @Input() message?: string;
  @Input() type: 'success' | 'info' | 'warning' | 'error' = 'info';
  @Output() closed = new EventEmitter<void>();

  get typeClass(): string {
    const classMappings = {
      success: 'bi-check-circle-fill text-appSuccess',
      info: 'bi-info-circle-fill text-appPrimary',
      warning: 'bi-exclamation-diamond-fill text-appWarning',
      error: 'bi-x-circle-fill text-appError',
    };

    return classMappings[this.type] || '';
  }

  close(): void {
    this.closed.emit();
  }
}
