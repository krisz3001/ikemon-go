import { Injectable } from '@angular/core';
import { User } from '../interfaces/user.interface';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor() { }

  user: User = {
    UserId: "asd-123-asd",
    Username: "krisz",
    Email: "asd@asd.asd",
    Password: "hashedpw",
    Money: 1000, 
    Admin: false,
    Created: "2024-01-20T20-42-40"
  };

  GetUser() : User | null {
    return this.user;
  }

}
