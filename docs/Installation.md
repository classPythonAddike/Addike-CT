<div align="center">
  <h1>Addike-CT</h1>
</div>

## ⬇️ Download and Install

There are two ways to download Addike-CT

### From Releases

1. Head over to the [Releases tab](https://github.com/classPythonAddike/Challenge-Tester-Backend/releases), and download the zip file which matches your operating system. Unzip the file. You should see an executable, and a folder named `public` in the unzipped contents.
2. Move the executable into your system path.
3. Copy the `public` folder into a suitable location.
4. Finally, create an environment variable called `CTPath`. Set its value to the location of the `public` folder.

### From Source
1. Fork, and clone this repository to your computer.
2. Create an environment variable called `CTPath`. Set its value to a path where you would like to store Addike-CT's files.
3. Run `make target="path/to/install/addike-ct" install`
4. An executable will be built in the folder provided. Make sure that it is a part of the system path.

#### Finally, verify your installation by typing `addike-ct --version` in the terminal. You should receive something like this:
```
$ addike-ct --version
Addike-ct v0.1.0 (stable)
```
You may have to restart your terminal for these changes to take effect.

Once you're done, head over to [Setup](/Setup) to create your first suite of tests!