<div align="center">
    <img src="./assets/logo.png" width="150" />
    <h3>A simple tool to take screenshots using your terminal</h3>
</div>

&nbsp;
## ❗️ Install:
Arch based:
```bash
git clone https://aur.archlinux.org/ghost-git.git
cd ghost-git
makepkg -si
```

OR with a AUR Helper:
```bash
yay -S ghost-git
```

&nbsp;
Others:
```bash
git clone https://github.com/z3oxs/ghost.git
cd ghost
sudo make install
```

&nbsp;
## 🚀 Usage:
**If you want to use the properly builded script, check <a href="#setup">setup</a>**

| Command    |     Description    |
| ------------- | ------------------ |
| f |  Take fullscreen shot (include all monitors) |
| s | Take screenshot from a selected area |
| g | Take screenshot from a selected area freezed and support multiple displays |
| d | Take screenshot from only a display |
| S | Save the output to a image file |
| c | Copy the output to the clipboard |
| o | Outputs screenshot to stdout encoded as png |
| u | Upload the screenshot to AnonFiles |
| v | Show version |

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
