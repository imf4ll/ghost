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
| version | Show version |

&nbsp;
### Example
#### Fullscreen and save to a png file inside your working folder
```bash
ghost -f -S 'image.png'
```
#### Selected area and copy to clipboard
```bash
ghost -s -c
```
#### Fullscreen and output to stdout
```bash
ghost -f -o
```

<br><br>
## ❌ Troubleshooting:
- If you're facing any "cc1.exe" error including "64-bit not implemented" (Windows):
    - Download [mingw-w64](https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/Personal%20Builds/mingw-builds/8.1.0/threads-posix/seh/x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z)
    - Extract and copy "mingw64" to "C:\"
    - Define or create "PATH" enviroment variable containing "C:\\mingw64\\bin\\"

<br><br>
<a name="setup"></a>
## 🔧 Setup:
### Clone this repository:<br>
`git clone https://github.com/z3oxs/ghost` or Download ZIP and unzip;<br><br>
### Move to repository:<br>
`cd ghost`<br><br>
### Install (Will need 'make' or download prebuilt binaries on releases page):<br>
#### Windows:<br>
`sudo make install-win`<br><br>
#### Linux:<br>
`sudo make install-linux`<br><br>
