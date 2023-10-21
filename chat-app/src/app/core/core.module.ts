import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MainComponent } from './layouts/main/main.component';
import { NavbarComponent } from './layouts/navbar/navbar.component';
import { SpinnerComponent } from './layouts/spinner/spinner.component';
import { RouterModule } from '@angular/router';
import { HasAnyAuthorityDirective } from './directives/has-any-authority.directive';

@NgModule({
  declarations: [MainComponent, NavbarComponent, SpinnerComponent, HasAnyAuthorityDirective],
  imports: [CommonModule, RouterModule],
  exports: [MainComponent, HasAnyAuthorityDirective],
})
export class CoreModule {}
