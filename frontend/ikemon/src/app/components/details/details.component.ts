import { Component, Input, OnChanges, OnInit } from '@angular/core';
import { Card } from '../../interfaces/card.interface';
import { ActivatedRoute, Router } from '@angular/router';
import { CardService } from '../../services/card.service';

@Component({
  selector: 'app-details',
  standalone: true,
  imports: [],
  templateUrl: './details.component.html',
  styleUrl: './details.component.scss'
})
export class DetailsComponent implements OnInit {
  constructor(private route: ActivatedRoute, private router: Router, private cardService: CardService) {}

  card: Card = {
    attack: 0,
    defense: 0,
    description: "",
    hp: 0,
    id: "",
    image: "",
    islocked: false,
    name: "",
    ownerid: "",
    price: 0,
    typeid: ""
  };

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      let id = Array.isArray(params["id"]) ? params["id"][0] : params["id"];
      
      this.cardService.GetCardById(id).subscribe(card => {
        if (Object.keys(card).length == 0) {
          this.router.navigate(['/']);

        }
        this.card = card!;
      })
    })
    
  }
}
