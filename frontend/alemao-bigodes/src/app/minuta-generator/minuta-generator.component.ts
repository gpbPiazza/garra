import { Clipboard } from '@angular/cdk/clipboard';
import { CommonModule } from '@angular/common';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';

@Component({
  selector: 'app-minuta-generator',
  standalone: true,
  imports: [
    CommonModule, 
    HttpClientModule, 
    MatButtonModule, 
    MatCardModule, 
    MatIconModule, 
    MatProgressSpinnerModule,
    MatSnackBarModule,
    MatTooltipModule
  ],
  template: `
    <div class="minuta-container">
      <mat-card class="minuta-card">
        <mat-card-header>
          <mat-card-title>Gerar minuta!</mat-card-title>
          <mat-card-subtitle>Selecione um arquivo ato consultar e gere sua minuta!</mat-card-subtitle>
        </mat-card-header>
        
        <mat-card-content>
          <button type="button" mat-raised-button *ngIf="!selectedFile" class="file-upload-container">
            <label for="pdf-upload" class="file-selector">
              Selecione arquivo PDF
              <input 
                type="file" 
                id="pdf-upload" 
                accept="application/pdf"
                (change)="onFileSelected($event)"
                class="file-input"
              >
            </label>
          </button>
          
          <div *ngIf="selectedFile" class="selected-file-container">
            <div class="selected-file-info">
              <mat-icon>description</mat-icon>
              <span class="file-name">{{ selectedFile.name }}</span>
              <button 
                mat-icon-button 
                color="warn" 
                class="remove-file-button"
                (click)="clearSelectedFile()"
                matTooltip="Remover arquivo">
                <mat-icon>close</mat-icon>
              </button>
            </div>
            <button 
              mat-raised-button 
              color="primary" 
              class="generate-button"
              (click)="generateMinuta()">
              GERAR MINUTA
            </button>
          </div>
        </mat-card-content>
      </mat-card>
      
      <div class="result-container" *ngIf="minutaResult">
        <mat-card class="result-card">
          <div class="card-header-actions">
            <mat-card-header>
              <mat-card-title>Minuta Gerada</mat-card-title>
            </mat-card-header>
            <button 
              mat-icon-button 
              color="primary" 
              class="copy-button" 
              (click)="copyMinutaContent()"
              matTooltip="Copiar Conteúdo">
              <mat-icon>content_copy</mat-icon>
            </button>
          </div>
          <mat-card-content>
            <div [innerHTML]="minutaResult"></div>
          </mat-card-content>
        </mat-card>
      </div>
      
      <div class="loading-spinner" *ngIf="isLoading">
        <mat-spinner></mat-spinner>
      </div>
    </div>
  `,
  styles: [` 
    .minuta-container {
      display: flex;
      flex-direction: column;
      align-items: center;
    }
    
    .minuta-card {
      width: 100%;
      max-width: 600px;
    }
    
    mat-card-header {
      justify-content: center;
    }
    
    mat-card-title {
      text-align: center;
      font-weight: 600;
      font-size: 42px;
      line-height: 52px;
    }

    mat-card-subtitle {
      max-width: 800px;
      margin: 8px auto 0;
      line-height: 32px;
      font-size: 22px;
      font-weight: 400;
    }
    
    .file-selector{
      display: flex;
      justify-content: center;
      align-items: center;    

      min-width: 330px;
      
      padding: 24px 48px;
      font-weight: 500;
      font-size: 24px;
      line-height: 28px;
      cursor: pointer;
    }
    

    .file-upload-container {
      display: flex;
      justify-content: center;
      margin: 1rem 0;
    }
    
    .file-input {
      display: none;
    }
    
    .selected-file-container {
      display: flex;
      flex-direction: column;
      align-items: center;
      width: 100%;
      margin: 1.5rem 0;
      gap: 1.5rem;
    }
    
    .selected-file-info {
      display: flex;
      align-items: center;
      background-color: rgba(0, 0, 0, 0.04);
      padding: 12px 16px;
      border-radius: 8px;
      width: 80%;
      max-width: 400px;
    }
    
    .selected-file-info mat-icon {
      margin-right: 12px;
      color: #e5322d;
    }
    
    .file-name {
      flex: 1;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-size: 18px;
    }
    
    .remove-file-button {
      margin-left: 8px;
    }
    
    .generate-button {
      font-size: 18px;
      min-width: 200px;
      padding: 8px 24px;
    }
    
    .result-container {
      margin-top: 2rem;
      width: 100%;
      max-width: 800px;
    }
    
    .result-card {
      width: 100%;
      box-shadow: 0px 2px 4px -1px rgba(0, 0, 0, 0.2), 
                  0px 4px 5px 0px rgba(0, 0, 0, 0.14), 
                  0px 1px 10px 0px rgba(0, 0, 0, 0.12);
    }

    .card-header-actions {
      display: flex;
      justify-content: space-between;
      align-items: center;
      width: 100%;
      padding-right: 16px;
    }
    
    .copy-button {
      margin-top: 8px;
    }
    
    .loading-spinner {
      display: flex;
      justify-content: center;
      margin-top: 2rem;
    }
  `]
})
export class MinutaGeneratorComponent {
  selectedFile: File | null = null;
  minutaResult: SafeHtml | null = null;
  rawHtmlContent: string = '';
  isLoading = false;
  
  constructor(
    private http: HttpClient, 
    private sanitizer: DomSanitizer,
    private clipboard: Clipboard,
    private snackBar: MatSnackBar
  ) {}
  
  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
    }
  }
  
  clearSelectedFile() {
    this.selectedFile = null;
    this.clearMinutaResult();
  }

  clearMinutaResult() {
    this.minutaResult = null;
    this.rawHtmlContent = '';
  }
  
  generateMinuta() {
    if (!this.selectedFile) return;
    
    this.isLoading = true;
    this.minutaResult = null;
    
    const formData = new FormData();
    formData.append('ato_consultar_pdf', this.selectedFile);
    
    this.http.post('http://localhost:8080/api/v1/generator/minuta', formData, {
      responseType: 'text'
    })
      .subscribe({
        next: (htmlContent) => {
          console.log('Minuta generated successfully');
          this.rawHtmlContent = htmlContent;
          this.minutaResult = this.sanitizer.bypassSecurityTrustHtml(htmlContent);
          this.isLoading = false;
        },
        error: (error) => {
          console.error('Error generating minuta:', error);
          this.isLoading = false;
          this.snackBar.open('Erro ao gerar minuta. Tente novamente.', 'Fechar', {
            duration: 5000,
            panelClass: ['error-snackbar']
          });
        }
      });
  }

  copyMinutaContent() {
    if (this.rawHtmlContent) {
      this.clipboard.copy(this.rawHtmlContent);
      this.snackBar.open('Conteúdo copiado para a área de transferência', 'Fechar', {
        duration: 3000,
        horizontalPosition: 'center',
        verticalPosition: 'bottom'
      });
    }
  }
}