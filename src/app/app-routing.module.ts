import { ValidatorsComponent } from './validator/validators/validators.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { AlarmsComponent } from './alarm/alarms/alarms.component'; 
import { ValidatorEditComponent } from './validator/validator-edit/validator-edit.component';
 
const routes: Routes = [
    { path: '', pathMatch: 'full', redirectTo: '/alarm/alarms' },
    { path: 'alarm/alarms', component: AlarmsComponent },
    { path: 'validator', component: ValidatorsComponent },
    { path: 'validator/edit', component: ValidatorEditComponent },
    { path: '**', component: PageNotFoundComponent },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
