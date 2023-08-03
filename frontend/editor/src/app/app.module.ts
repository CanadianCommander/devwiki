import {ApplicationRef, Inject, Injector, NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

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
