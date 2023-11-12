export class CustomButton extends HTMLElement {
  public static readonly componentName = "custom-button";

  private readonly element = document.createElement("button");
  private readonly shadow = this.attachShadow({ mode: "closed" });
  private class: string[] = [
    "px-4",
    "py-1",
    "text-zinc-50",
    "border",
    "border-zinc-50",
    "rounded",
    "hover:bg-zinc-50",
    "hover:text-zinc-800",
  ];

  constructor() {
    super();
    this.init();
    this.render();
  }

  protected init() {
    this.element.classList.add(...this.class);
    this.shadow.appendChild(this.element);
  }

  protected render() {
    this.element.innerHTML = `
      <span class="text-uppercase"> 
        Click me
      </span>
    `;
  }
}
