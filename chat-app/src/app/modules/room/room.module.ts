import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { RouterModule } from '@angular/router';
import { ComponentsModule } from 'src/app/components/components.module';
import { RoomComponent } from './room.component';
import { routes } from './room.routes';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [RoomComponent],
  imports: [CommonModule, RouterModule.forChild(routes), ComponentsModule, FormsModule],
  exports: [RoomComponent],
})
export class RoomModule {}
