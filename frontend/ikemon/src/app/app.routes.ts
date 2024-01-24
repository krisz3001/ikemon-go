import { Routes } from '@angular/router';
import { ModalComponent } from './components/modal/modal.component';
import { CardListComponent } from './components/card-list/card-list.component';
import { DetailsComponent } from './components/details/details.component';
import { detailsGuard } from './guards/details.guard';
import { AccountComponent } from './components/account/account.component';
import { accountGuard } from './guards/account.guard';

export const routes: Routes = [
    { path: '', component: CardListComponent},
    { path: 'details', component: DetailsComponent, canActivate: [detailsGuard]},
    { path: 'account', component: AccountComponent, canActivate: [accountGuard]},
    { path: 'login', component: ModalComponent},
    { path: '**', component: ModalComponent},
];
