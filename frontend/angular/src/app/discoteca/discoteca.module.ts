// Cosas de angular
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from "@angular/common";

import { DiscotecaComponent } from './component/discoteca.component';
import { DiscotecaDetailsComponent } from './component/discoteca-details.component';
import { DiscotecaPreviewComponent } from './component/discoteca-preview.component';
import { DiscotecaResolver } from './component/discoteca-resolver.service';
import { DiscotecaRoutingModule } from './discoteca-routing.module';

@NgModule({
  declarations: [
    DiscotecaComponent,
    DiscotecaDetailsComponent,
    DiscotecaPreviewComponent
  ],
  imports: [
    DiscotecaRoutingModule,
    CommonModule
  ],
  providers: [
    DiscotecaResolver
  ]
})

export class DiscotecaModule {}