import { CustomButton } from "./components/Button/Button";
import { EditorView } from "./components/Editor/EditorView";

function defineComponents() {
  window.customElements.define(CustomButton.componentName, CustomButton);
  window.customElements.define(EditorView.componentName, EditorView);
}

function main() {
  defineComponents();
}

main();
