import {ApplicationRef, Inject, Injector, NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {EditorComponent} from "./editor.component";
import {createCustomElement} from "@angular/elements";

@NgModule({
  declarations: [
    EditorComponent
  ],
  exports: [
      EditorComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
})
export class AppModule {

  // =======================================================================
  // Public
  // =======================================================================

  constructor(private injector: Injector) {}

  // =======================================================================
  // Overrides
  // =======================================================================

  public ngDoBootstrap():void {
    const newEl = createCustomElement(EditorComponent, {injector: this.injector});
    customElements.define('wiki-editor', newEl);
  }

}
