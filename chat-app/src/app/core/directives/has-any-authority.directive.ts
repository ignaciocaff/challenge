import { Directive, Input, OnDestroy, TemplateRef, ViewContainerRef } from '@angular/core';
import { Subject, map, takeUntil } from 'rxjs';
import { SessionService } from '../services/session.service';

/**
 * @whatItDoes Conditionally includes an HTML element if current user has any
 * of the authorities passed as the `expression`.
 *
 * @howToUse
 * ```
 *     <some-element *appHasAnyAuthority="'ROLE_ADMIN'">...</some-element>
 *     <some-element *appHasAnyAuthority="['ROLE_ADMIN', 'ROLE_USER']">...</some-element>
 * ```
 */
@Directive({
  selector: '[appHasAnyAuthority]',
})
export class HasAnyAuthorityDirective implements OnDestroy {
  private authorities!: string | string[];
  private readonly destroy$ = new Subject<void>();

  constructor(
    private templateRef: TemplateRef<unknown>,
    private viewContainerRef: ViewContainerRef,
    private readonly sessionService: SessionService
  ) {}

  @Input()
  set appHasAnyAuthority(value: string | string[]) {
    this.authorities = value;
    this.updateView();

    // Get notified each time authentication state changes.
    this.sessionService
      .getUserRoles()
      .pipe(takeUntil(this.destroy$))
      .subscribe(() => {
        this.updateView();
      });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  private updateView(): void {
    this.sessionService
      .getUserRoles()
      .pipe(
        map((roles) => roles.some((role) => this.authorities.includes(role.name))),
        takeUntil(this.destroy$)
      )
      .subscribe((hasAnyAuthority) => {
        this.viewContainerRef.clear();
        if (hasAnyAuthority) {
          this.viewContainerRef.createEmbeddedView(this.templateRef);
        }
      });
  }
}
