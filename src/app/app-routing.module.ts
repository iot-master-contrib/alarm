import { ValidatorsComponent } from './validator/validators/validators.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AlarmsComponent } from './alarm/alarms/alarms.component';
import { ValidatorEditComponent } from './validator/validator-edit/validator-edit.component';
import {PageNotFoundComponent} from "./base/page-not-found/page-not-found.component";

const routes: Routes = [
    { path: '', pathMatch: 'full', redirectTo: 'alarm' },
    { path: 'alarm', component: AlarmsComponent },
    { path: 'validator', component: ValidatorsComponent },
    { path: 'validator/edit/:id', component: ValidatorEditComponent },
    { path: 'validator/create', component: ValidatorEditComponent },
    { path: '**', component: PageNotFoundComponent },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
