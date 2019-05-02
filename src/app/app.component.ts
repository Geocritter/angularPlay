import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Tour of Heroes';
  playAudio() {
    let audio = new Audio();
    audio.src = "assets/proofofhero.mp3";
    audio.load();
    audio.play();
  }
  ngOnInit() {
  }
}
