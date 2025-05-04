import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';


@Component({
  selector: 'app-toolbar',
  imports: [MatToolbarModule, MatButtonModule, MatIconModule],
  template: `
    <header>
      <mat-toolbar role="heading">
        <span>Notas</span>
      </mat-toolbar>
    </header>
  `,
  styles: ``
})
export class ToolbarComponent {

}
