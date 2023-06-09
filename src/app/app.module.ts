import { AlarmsComponent } from './alarm/alarms/alarms.component';

import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import {APP_BASE_HREF, registerLocaleData} from '@angular/common';
import zh from '@angular/common/locales/zh';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzModalModule } from 'ng-zorro-antd/modal';
import {NzCardModule} from "ng-zorro-antd/card"; 
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzInputModule } from 'ng-zorro-antd/input';
import { ValidatorsComponent } from './validator/validators/validators.component';
import { ValidatorEditComponent } from './validator/validator-edit/validator-edit.component';
import {BaseModule} from "./base/base.module";
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzTagModule } from 'ng-zorro-antd/tag';
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    AlarmsComponent,  ValidatorsComponent, ValidatorEditComponent
  ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        NzInputModule,
        NzMessageModule,
        NzButtonModule,
        NzTableModule,
        ReactiveFormsModule,
        NzFormModule ,
        NzModalModule,
        NzTagModule,
        NzCardModule,
        NzInputNumberModule,
        NzIconModule,
        NzSelectModule,
        NzSpaceModule,
        NzSpinModule,
        FormsModule,
        NzDividerModule,
        HttpClientModule,
        BrowserAnimationsModule,
        BaseModule
    ],
  providers: [
    { provide: APP_BASE_HREF, useValue: '/app/alarm/' },
    { provide: NZ_I18N, useValue: zh_CN }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
