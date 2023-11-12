import { CustomButton } from "./components/Button/Button";

function defineComponents() {
  window.customElements.define(CustomButton.componentName, CustomButton);
}

function main() {
  defineComponents();
}

main();
