import tailwind from "../../base.css?inline";

export class CustomButton extends HTMLElement {
  public static readonly componentName = "custom-button";
  public static readonly extends = "button";

  private readonly node = document.createElement("button");
  private readonly styleNode = document.createElement("style");
  private readonly shadow = this.attachShadow({ mode: "closed" });
  private class: string[] = [
    "px-4",
    "py-1",
    "text-zinc-700",
    "border",
    "border-zinc-50",
    "rounded",
    "hover:bg-zinc-50",
    "hover:text-zinc-700",
  ];

  constructor() {
    super();
    this.render();
    this.init();
  }

  public connectedCallback() {
    // event listeners
    this.node.addEventListener("click", this.onClick);
  }

  public disconnectedCallback() {
    // remove event listeners
    this.node.removeEventListener("click", this.onClick);
  }

  protected init() {
    this.node.classList.add(...this.class);
    this.styleNode.textContent = tailwind;
    this.shadow.appendChild(this.styleNode);
    this.shadow.appendChild(this.node);
  }

  protected render() {
    this.node.innerHTML = `
      <span class="text-uppercase"> 
        <slot></slot>
      </span>
    `;
  }

  private onClick() {
    console.log("!");
  }
}
