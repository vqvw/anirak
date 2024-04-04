<p align="center">
    <img src="https://i.imgur.com/y6exXqM.png" width="100"/>
    <br>
    <a style="font-size:larger" href="https://github.com/vqvw/anirak">Anirak Theme for Visual Studio Code</a>
    <br>
    <span>Supported Languages:</span>
    <i>C, C++, C#, CSS, SCSS, Go, HTML, Java, JS, JSX, JSON, Lua, Perl,<br>PHP, PHTML, Prisma, Python, Ruby, Rust, Shell, SQL, Swift, TypeScript, TSX, YAML</i>
    <br />
    <br />
    <img src="https://i.imgur.com/HvwwQeM.png" width="500"/>
</p>

## Maintenance

After 5 years of active maintenance, the Anirak theme is being sunset. The
theme is still in a working state and should continue to function as normal for
some time. New features will not be supported.

## Installation

Using the VS Code command palette: `ext install barjo.anirak`

Or from the <a href="https://marketplace.visualstudio.com/items?itemName=barjo.anirak">Visual Studio Marketplace</a>

## Screenshots

![](https://i.imgur.com/iGz4YCH.png)

![](https://i.imgur.com/cFQAl8Q.png)

![](https://i.imgur.com/vo6aZca.png)

![](https://i.imgur.com/RayB15j.png)

## Settings

**Note:** Using a font which has a clear **bold** variant is useful when using this theme. The font used in the screenshots is [Roboto Mono](https://fonts.google.com/specimen/Roboto+Mono).

```json
{
  "editor.fontFamily": "Roboto Mono",
  "editor.fontSize": 15,
  "editor.lineHeight": 30,
  "editor.letterSpacing": -0.75
}
```

## Description

Anirak is a modern dark/light theme, aiming to combine practicality and elegance with simplicity. The theme is minimalistic, using mostly blueish whites and greys, with carefully chosen accents of vibrant blue, light blue, cyan, and green. An assortment of other colours are used subtley throughout the theme where needed. Anirak follows a philosophy of using colour to bring attention to pertinent pieces of code, leaving the rest uncoloured. The theme makes use of bold and italic text to further increase the readablity of uncoloured code.

As a general rule of thumb, code is coloured as follows:

| Colour           | Used for                                                                                    |
| ---------------- | ------------------------------------------------------------------------------------------- |
| Vibrant blue     | Strings                                                                                     |
| Light blue       | Data-types and some keywords (var, function, package, import, etc.)                         |
| Cyan             | Storage modifiers and special elements (public, private, static, as well as JSX components) |
| Vibrant green    | Numbers and constants (true/false, null, "\n", "\e", etc.)                                  |
| None             | Default text colour and variables                                                           |
| None bold        | Functions                                                                                   |
| None italic      | Object properties                                                                           |
| Grey             | Punctuation and some keywords (return, if, for, while, struct, new, etc.)                   |
| Dark grey italic | Comments                                                                                    |
