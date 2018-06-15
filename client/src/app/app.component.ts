import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { isArray } from 'util';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Boggle Solver';
  rows = 4;
  cols = 4;
  lang = 'en_US';
  board: any = {};
  words: string[];
  loading = false;
  constructor(private http: HttpClient) {

  }

  public rowsUpdate($event: any) {
    this.rows = $event.target.value;
  }
  public colsUpdate($event: any) {
    this.cols = $event.target.value;
  }
  public langUpdate($event: any) {
    this.lang = $event.target.value;
  }
  public generateBoard() {
    const acceptedChars = new RegExp(/^[A-Za-z]+$/);
    this.board = {
      lang: this.lang,
      rows: []
    };

    for (let ri = 0; ri < this.rows; ri++) {
      const row = {
        cols: []
      };
      this.board.rows.push(row);
      for (let ci = 0; ci < this.cols; ci++) {
        let ch = String.fromCharCode(Math.floor(Math.random() * 256)).toLowerCase();
        while (!acceptedChars.test(ch)) {
          ch = String.fromCharCode(Math.floor(Math.random() * 256)).toLowerCase();
        }
        row.cols.push({
          char: ch
        });
      }
    }

    this.words = [];
    this.loading = false;

  }

  public showPossibleWords() {
    this.loading = true;
    this.words = [];
    this.http.post('/api/possiblewords', this.board).subscribe((data: string[]) => {
      this.words = data;
      this.loading = false;
    });
  }

}
