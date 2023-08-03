import {Component, Injector} from '@angular/core';
import {EditorComponent} from "./editor.component";
import {createCustomElement} from "@angular/elements";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  constructor(injector: Injector) {
    createCustomElement(EditorComponent, {injector: injector});
  }

}
