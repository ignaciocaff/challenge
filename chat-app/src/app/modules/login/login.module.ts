import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { RouterModule } from '@angular/router';
import { ComponentsModule } from 'src/app/components/components.module';
import { LoginComponent } from './login.component';
import { routes } from './login.routes';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [LoginComponent],
  imports: [CommonModule, RouterModule.forChild(routes), ComponentsModule, FormsModule],
  exports: [LoginComponent],
})
export class LoginModule {}
