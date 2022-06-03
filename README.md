<div align="center">
    <img src="./assets/logo.png" width="150" />
    <h3>A simple script to take screenshots using your terminal</h3>
</div>

&nbsp;
## 🚀 Usage:
**If you want to use the properly builded script, check <a href="#setup">setup</a>**

| Command    |     Description    |
| ------------- | ------------------ |
| fullscreen |  Take fullscreen shot (include all monitors) |
| selection | Take screenshot from a selected area |
| selectiongui | Take screenshot from a selected area freezed and support multiple displays |
| display | Take screenshot from only a display |
| save | Save the output to a image file |
| clipboard | Copy the output to the clipboard |
| output | Outputs screenshot to stdout encoded as png |
| upload | Upload the screenshot to AnonFiles |

&nbsp;
### Example
#### Fullscreen and save to a png file inside your working folder
```bash
goshot -f -S 'image.png'
```
#### Selected area and copy to clipboard
```bash
goshot -s -c
```
#### Fullscreen and output to stdout
```bash
goshot -f -o
```

<br><br>
<a name="setup"></a>
## 🔧 Setup:
### Clone this repository:<br>
`git clone https://github.com/z3oxs/goshot` or Download ZIP and unzip;<br><br>
### Move to repository:<br>
`cd goshot`<br><br>
### Install (Will need 'make'):<br>
`sudo make install`<br><br>
