import tailwind from "../../base.css?inline";

export class EditorView extends HTMLElement {
  public static readonly componentName = "editor-view";

  private readonly node = document.createElement("html");
  private readonly styleNode = document.createElement("style");
  private shadow = this.attachShadow({ mode: "closed" });

  constructor() {
    super();
    this.init();
    this.render();
  }

  public init() {
    this.styleNode.textContent = tailwind;
    this.shadow.appendChild(this.styleNode);
    this.shadow.appendChild(this.node);
  }

  public render() {
    this.node.innerHTML = `
      <head>
        <title>Your shid</title>
      </head> 

      <body>
        <h1>Oh shidd</h1>
      </body>
    `;
  }
}
