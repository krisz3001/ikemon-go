import { Injectable } from '@angular/core';
import { Card } from '../interfaces/card.interface';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CardsResponse } from '../interfaces/cardresponse.interface';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class CardService {

  constructor(private httpClient: HttpClient) {}

  GetCardById(id: string): Observable<any>{
    return this.httpClient.get(`${environment.backendURL}/cards`, {
      params: {
        id: id
      }, 
      withCredentials: true
    })
  }

  GetNthNineOfType(n: number, type: string, userId: string = ""): Observable<CardsResponse> {
    return this.httpClient.get<CardsResponse>(`${environment.backendURL}/cards`, {
      params: {
        page: n,
        type: type
      },
      withCredentials: true
    });
  }
}