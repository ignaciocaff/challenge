import { Component } from '@angular/core';
import { Breadcrumb } from './models/Breadcrumb';
import { Observable } from 'rxjs';
import { BreadcrumbsService } from './breadcrumbs.service';

@Component({
  selector: 'app-breadcrumbs',
  templateUrl: './breadcrumbs.component.html',
})
export class BreadcrumbsComponent {
  public breadcrumbs$: Observable<Breadcrumb[]>;

  constructor(private breadcrumbsService: BreadcrumbsService) {
    this.breadcrumbs$ = breadcrumbsService.breadcrumbs$;
  }
}
