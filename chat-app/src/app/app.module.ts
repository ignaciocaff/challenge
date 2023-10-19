import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { CoreModule } from './core/core.module';
import { MainComponent } from './core/layouts/main/main.component';
import { HttpClientModule } from '@angular/common/http';
import { httpInterceptorProviders } from './core/interceptors';
import { ComponentsModule } from './components/components.module';
import { BrowserModule } from '@angular/platform-browser';

@NgModule({
  declarations: [
  ],
  imports: [BrowserModule, HttpClientModule, AppRoutingModule, CoreModule, ComponentsModule],
  providers: [...httpInterceptorProviders],
  bootstrap: [MainComponent],
})
export class AppModule { }
