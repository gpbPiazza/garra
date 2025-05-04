import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MinutaGeneratorComponent } from './minuta-generator/minuta-generator.component';
import { ToolbarComponent } from './toolbar/toolbar.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule, 
    MatToolbarModule, 
    MatIconModule,
    FormsModule,
    MinutaGeneratorComponent,
    ToolbarComponent,
  ],
  template: `
    <div>
      <app-toolbar></app-toolbar>
      <div class="content">
        <app-minuta-generator></app-minuta-generator>
      </div>
    </div>
  `,
  styles: [`
    .content {
      padding-top: 60px; 
    }
  `],
})
export class AppComponent  {}
