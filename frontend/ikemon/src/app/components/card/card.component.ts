import { Component, Input, OnInit } from '@angular/core';
import { Card } from '../../interfaces/card.interface';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { CardType } from '../../interfaces/cardtype.interface';

@Component({
  selector: 'app-card',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './card.component.html',
  styleUrl: './card.component.scss'
})
export class CardComponent {

  constructor() {}

  @Input() card!: Card;

  types!: CardType[];

}
