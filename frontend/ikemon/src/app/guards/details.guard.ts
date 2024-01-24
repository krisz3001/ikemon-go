import { CanActivateFn, Router } from '@angular/router';
import { CardService } from '../services/card.service';
import { inject } from '@angular/core';
import { Card } from '../interfaces/card.interface';
import { Observable, of, switchMap } from 'rxjs';

export const detailsGuard: CanActivateFn = (route, state) => {
  let cardService: CardService = inject(CardService)
  let router: Router = inject(Router);

  /* return cardService.GetCardById(route.queryParams["id"]).pipe(
    switchMap(res => {
      if (Object.keys(res).length) {
        return of(true)
      }
      router.navigate(['/']);
      return of(false)
    })
  ); */
  return true
  
};