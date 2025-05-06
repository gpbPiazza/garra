import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { HeaderComponent } from './header/header.component';
import { MinutaGeneratorComponent } from './minuta-generator/minuta-generator.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    MatToolbarModule,
    MatIconModule,
    FormsModule,
    MatButtonModule,
    MinutaGeneratorComponent,
    HeaderComponent,
],
  template: `
    <app-header></app-header>
    <div class="content">
      <app-minuta-generator></app-minuta-generator>
    </div>
  `,
  styles: [`
    :host {
      display: flex;
      flex-direction: column;
      height: 100vh;
      overflow: hidden;
    }
    
    .content {
      padding-top: 60px; 
      background-color: var(--mat-sys-background);
      color: var(--mat-sys-on-background);
      height: 100vh;
      width: 100%;
    }
  `],
})
export class AppComponent  {}
