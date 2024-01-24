import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { UserService } from '../../services/user.service';
import { User } from '../../interfaces/user.interface';
import { CommonModule } from '@angular/common';
import { CardService } from '../../services/card.service';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent implements OnInit {
  constructor(private userService: UserService, private cardService: CardService, private router: Router) {}

  user!: User | null;

  cards_count: number = 0;
  
  ngOnInit(): void {
    this.user = this.userService.GetUser();
    
    /* if (this.user) {
      this.cards_count = this.cardService.GetOwnedCardsCount(this.user.UserId)
    } */
    
  }
}
