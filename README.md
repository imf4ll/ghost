<div align="center">
    <img src="./assets/logo.png" width="150" />
    <h3>A tool to take screenshots using your Linux terminal</h3>
</div>

<br><br>
## ❗️ Install:
```bash
git clone https://github.com/z3oxs/ghost.git
cd ghost
make install
```

&nbsp;
## 🚀 Usage:
**If you want to use the properly builded script, check <a href="https://github.com/z3oxs/ghost/releases">releases page</a>**

&nbsp;
<div align="center">
    
| Command    |     Description    |
| ------------- | ------------------ |
| f |  Take fullscreen shot (include all monitors) |
| s | Take screenshot from a selected area |
| g | Take screenshot from a selected area freezed and support multiple displays |
| d | Take screenshot from only a display |
| S | Save the output to a image file |
| F | Save to file on default images path |
| c | Copy the output to the clipboard |
| o | Outputs screenshot to stdout encoded as png |
| u | Upload the screenshot to AnonFiles |
| v | Show version |
    
</div>

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
