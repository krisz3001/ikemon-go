import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { CardComponent } from '../card/card.component';
import { CommonModule } from '@angular/common';
import { CardService } from '../../services/card.service';
import { Card } from '../../interfaces/card.interface';
import { TypeService } from '../../services/type.service';
import { CardType } from '../../interfaces/cardtype.interface';
import { PagerService } from '../../services/pager.service';
import { Session } from '../../interfaces/session.interface';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-card-list',
  standalone: true,
  imports: [CardComponent, CommonModule, FormsModule],
  templateUrl: './card-list.component.html',
  styleUrl: './card-list.component.scss'
})


export class CardListComponent implements OnInit {

  constructor(private cardService: CardService, private typeService: TypeService, private pagerService: PagerService) {}

  @Input() userId?: string;

  cards!: Card[];
  types!: CardType[];

  session!: Session;
  selectedType!: string;
  pages!: number[]

  ngOnInit(): void {
    this.session = this.pagerService.GetSession();
    this.GetTypes();
    this.selectedType = this.session.Type;
  }

  GetCards(): void {
    this.cardService.GetNthNineOfType(this.session.Page, this.session.Type, this.userId)
    .subscribe(res => {
      this.cards = res.cards;
      this.pages = Array(res.pages);
      this.cards?.map(card => {
        if(card.image.startsWith("img/")) card.image = "../../../assets/" + card.image;
        let type = this.types.find(t => t.typeid == card.typeid);
        if (type) card.typename = type.name;
        else card.typename = card.typeid;
      });
    });
  }

  GetTypes(): void {
    this.typeService.GetTypes()
    .subscribe(res => {
      this.types = res.types;
      this.GetCards();
    })
  }

  GetSelectedType(type: string): boolean {
    return type === this.session.Type;
  }

  SelectType(e: string) {
    this.pagerService.SetType(e);
    this.pagerService.SetPage(0);
    this.GetCards();
  }

  Active(n: number): string {
    return n === this.session.Page ? "pager-active" : "";
  }

  TurnPage(n: number) {
    if (n === this.session.Page) return;
    this.pagerService.SetPage(n);
    this.GetCards();
  }

}