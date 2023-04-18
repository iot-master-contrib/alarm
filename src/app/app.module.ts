import { AlarmsComponent } from './alarm/alarms/alarms.component';
 
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser'; 
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component'; 
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component'; 
import { NzTableModule } from 'ng-zorro-antd/table'; 
import { NzDividerModule } from 'ng-zorro-antd/divider'; 
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzButtonModule } from 'ng-zorro-antd/button'; 
import { NzInputModule } from 'ng-zorro-antd/input'; 
import { ValidatorsComponent } from './validator/validators/validators.component';
import { ValidatorEditComponent } from './validator/validator-edit/validator-edit.component';
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent, 
    PageNotFoundComponent, 
    AlarmsComponent,  ValidatorsComponent, ValidatorEditComponent
  ],
  imports: [
    BrowserModule, 
    AppRoutingModule,
    NzInputModule,
    NzMessageModule,
    NzButtonModule,
    NzTableModule,
    NzModalModule,
    NzIconModule ,
    NzSpaceModule,
    NzSpinModule,
    FormsModule,
    NzDividerModule,
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [
    { provide: NZ_I18N, useValue: zh_CN }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
