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
          <div class="file-upload-container">
            <label for="pdf-upload" class="file-upload-label" [class.has-file]="selectedFile">
              <div class="upload-icon">
                <mat-icon>cloud_upload</mat-icon>
              </div>
              <div class="upload-text">
                {{ selectedFile ? selectedFile.name : 'Selecione um arquivo PDF' }}
              </div>
              <input 
                matInput
                type="file" 
                id="pdf-upload" 
                accept="application/pdf"
                (change)="onFileSelected($event)"
                class="file-input"
              >
            </label>
          </div>
        </mat-card-content>
        
        <mat-card-actions>
          <button 
            mat-raised-button 
            color="primary" 
            [disabled]="!selectedFile"
            (click)="generateMinuta()">
            GERAR
          </button>
        </mat-card-actions>
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
    }
    
    .file-upload-container {
      display: flex;
      justify-content: space-between;
      margin: 1rem 0;
    }
    
    .file-upload-label {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      border: 2px dashed #ccc;
      border-radius: 8px;
      padding: 2rem;
      width: 100%;
      cursor: pointer;
      transition: all 0.3s ease;
    }
  
    
    .upload-icon {
      font-size: 2rem;
      margin-bottom: 1rem;
    }
    
    .upload-icon mat-icon {
      font-size: 48px;
      height: 48px;
      width: 48px;
    }
    
    .upload-text {
      text-align: center;
      overflow: hidden;
      text-overflow: ellipsis;
      max-width: 100%;
    }
    
    .file-input {
      display: none;
    }
    
    mat-card-actions {
      display: flex;
      justify-content: center;
      padding-bottom: 1.5rem;
    }
    
    .result-container {
      margin-top: 2rem;
      width: 100%;
      max-width: 800px;
    }
    
    .result-card {
      width: 100%;
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