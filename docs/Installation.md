<div align="center">
  <h1>Addike-CT</h1>
</div>

## ⬇️ Download 

There are two ways to download Addike-CT

1. From Releases - Head over to the [Releases tab](https://github.com/classPythonAddike/Challenge-Tester-Backend/releases), and download the zip file which matches your operating system. Unzip the file. You should see an executable, and a folder named `public` in the unzipped contents.
2. From Source - Fork, and clone this repository to your computer. Navigate into the directory, then run `go build .` In addition to this, you will need to download the the minified files for the frontend of the tester. The full instructions for the same can be found [here](https://github.com/classPythonAddike/Challenge-Tester-Frontend).

## 📁 Install 

To install Addike-CT, make sure you have both the frontend (html + js + css files), and the backend (executable) on your machine.

1. Place the executable in a convenient location, for example - `C:\Users\Username\addike-ct\addike-ct.exe` on Windows. Then add this location to your system path.
2. Next, place the `public` folder in a convenient location, for example - `C:\Users\Username\addike-ct\public` on Windows. Then create a new environment variable by the name of `CTPath` and set it to the location of the public folder. In this example, it would be `CTPath='C:\Users\Username\addike-ct\public'`
3. Finally, verify your manual installation by typing `addike-ct --version` in the terminal. You should receive something like this:
```
$ addike-ct --version
Addike-ct v0.1.0 (beta)
  ```
You may have to restart your terminal for these changes to take effect.

Once you're done, head over to [Setup](/Setup) to create your first suite of tests!
