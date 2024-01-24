import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { TypeResponse } from '../interfaces/typeresponse.interface';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class TypeService {

  constructor(private http: HttpClient) { }

  GetTypes(): Observable<TypeResponse> {
    return this.http.get<TypeResponse>(`${environment.backendURL}/types`, {
      withCredentials: true
    });
  }
}
