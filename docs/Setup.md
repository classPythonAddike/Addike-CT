<div align="center">
  <h1> Addike-CT </h1>
</div>

Make sure you have installed, and verified your installation of Addike-CT before continuing

## ⛏️ Setup a Project to Test

1. Setup has been made much easier since the first stable release of Addike-CT. You can just run the following command to create a new project -

```sh
$ addike-ct --create
```

to set up a new project. You can then replace the default test files with your own.

Next, you'll need to configure Addike-CT for your test suite. In the `config/` directory, you should see a couple of files, but only three are important to us -
1. `cases.json` which holds the test cases.
2. `challenge-info.json` which describes the challenge, and provides the time limit for each test file.
3. `language-config.json` which lists the allowed languages, for this test suite.

You may change the values in these files, to suit your needs. You can additionally see and example of a script to generate random test cases in `config/generator.py`.

For example, the language-config.json file may look like this -
```json
{
    "Languages": [
        {
            "Language": "C", // The name of the language
            "Command": "g++ -o '[n]-c' '[f]'", // Command used to compile. Leave it blank if no compilation is required.
            "Execution": " ./[n]-c", // Command used to run.
			//  Use [n] to denote the filename, without extension, and [f] to denote the full filename
            "Extension": ".c", // File extension for the language.
            "CompiledLanguage": true
        },
		....
	]
}
```

## 📚 Running Tests

Once you have set up your project, you're now ready to start testing! Simply run the following command in your terminal -
```sh
$ addike-ct
```
Then navigate to `http://localhost:8080` to view the progress of your tests!

As of now, Addike-CT runs the tests locally, which can be a security issue if the code you're testing is from third parties. Don't worry too much though! We expect to release remote testing on WandBox soon!
