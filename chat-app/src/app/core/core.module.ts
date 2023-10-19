import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FooterComponent } from './layouts/footer/footer.component';
import { MainComponent } from './layouts/main/main.component';
import { NavbarComponent } from './layouts/navbar/navbar.component';
import { SpinnerComponent } from './layouts/spinner/spinner.component';
import { RouterModule } from '@angular/router';
import { BreadcrumbsComponent } from './layouts/breadcrumbs/breadcrumbs.component';
import { HasAnyAuthorityDirective } from './directives/has-any-authority.directive';
import { ToastContainerComponent } from './layouts/toast/components/container/toast-container.component';
import { ToastComponent } from './layouts/toast/components/toast.component';

@NgModule({
  declarations: [FooterComponent, MainComponent, NavbarComponent, SpinnerComponent, BreadcrumbsComponent, HasAnyAuthorityDirective, ToastContainerComponent, ToastComponent],
  imports: [CommonModule, RouterModule],
  exports: [MainComponent, HasAnyAuthorityDirective],
})
export class CoreModule {}
