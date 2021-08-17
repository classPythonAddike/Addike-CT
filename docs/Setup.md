<div align="center">
  <h1> Addike-CT </h1>
</div>

Make sure you have installed, and verified your installation of Addike-CT before continuing

## ⛏️ Setup 

1. Setup has been made much easier in the second release of addike-ct. You can just run

```sh
$ addike-ct --create
```

to set up a new project. You can then replace the default test files with your own, and fill out the data in `config/` with your own, and you should be good to go!

## 📚 Usage 

Once you have set up your project, you're now ready to start testing! There are two ways in which you can test files -
```sh
$ addike-ct
```
Or
```sh
$ addike-ct --validate
```
The latter tests files _only after validating them_. If you use this, you will need to have a `validate.py` in your `config` directory. This is generally recommended when testing third party code.
