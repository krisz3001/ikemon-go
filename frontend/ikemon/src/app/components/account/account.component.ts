import { Component, OnInit } from '@angular/core';
import { CardListComponent } from '../card-list/card-list.component';
import { UserService } from '../../services/user.service';
import { CardService } from '../../services/card.service';
import { Card } from '../../interfaces/card.interface';
import { User } from '../../interfaces/user.interface';

@Component({
  selector: 'app-account',
  standalone: true,
  imports: [CardListComponent],
  templateUrl: './account.component.html',
  styleUrl: './account.component.scss'
})
export class AccountComponent implements OnInit {
  constructor(private userService: UserService, private cardService: CardService) {}

  cards!: Card[];
  user!: User;
  cards_count: number = 0;

  ngOnInit(): void {
    this.user = this.userService.GetUser()!;
    /* this.cards = this.cardService.GetOwnedCards(this.user.UserId);
    this.cards_count = this.cardService.GetOwnedCardsCount(this.user.UserId); */
  }
}
