import { Injectable } from '@angular/core';
import { Session } from '../interfaces/session.interface';

@Injectable({
  providedIn: 'root'
})
export class PagerService {

  constructor() { }

  session: Session = {
    SessionID: "asdfgh0",
    Page: 0,
    Type: "0"
  }

  GetSession(): Session {
    return this.session;
  }

  SetPage(n: number) {
    this.session.Page = n;
  }

  SetType(type: string) {
    this.session.Type = type;
  }
}
