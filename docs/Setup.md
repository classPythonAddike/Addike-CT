<div align="center">
  <h1> Addike-CT </h1>
</div>

Make sure you have installed, and verified your installation of Addike-CT before continuing

## ⛏️ Setup a Project to Test

1. Setup has been made much easier since the first stable release of Addike-CT. You can just run the following command to create a new project -

```sh
$ addike-ct --create
```

to set up a new project. You can then replace the default test files with your own, and fill out the data in `config/` with your own, and you should be good to go!

## 📚 Running Tests

Once you have set up your project, you're now ready to start testing! Simply run the following command in your terminal -
```sh
$ addike-ct
```
Then navigate to `http://localhost:8080` to view the progress of your tests!

As of now, Addike-CT runs the tests locally, which can be a security issue if the code you're testing is from third parties. Don't worry too much though! We expect to release remote testing on WandBox soon!