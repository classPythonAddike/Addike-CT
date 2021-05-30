## ⛏️ Setup 

1. In your workspace, where your Python files are, create a folder named `config`.

2. Next, create a file in `config` with the name `challenge-info.json`. This file will be used to provide the tester with information about the testing. Open the file, and paste the following text in it:
```json
{
  "ChallengeName": <Title>,
  "ChallengeDescription": <Description>,
  "CleanUpCommand": <Command>,
  "Timeout": <Timeout>
}
```
Where:
```md
1. Title (string) - The name of the challenge
2. Description (string) - The challenge's description
3. Command (string) - A command to run, after the testing has finished. If you don't want to run anything after the testing, set it to something like `````cls```
4. Timeout (integer) - The maximum time allowed, in seconds for a file to finish all tests
```

3. (Optional) Create a file named `validate.py` in the `config` directory. This will be used to check whether a file is valid, to ensure that certain conditions are met (like no imports, eval, exec, etc). The file that `validate.py` will have to check will be passed in as a command line argument. You can see a sample of this over [here](./pythontests/config/validate.py)

4. Finally, you need to create some testcases! Make a file named `cases.json` in the `config` directory. In it, you can write the input and output for each test, like so -

```json
{
  "Cases": [
  {
    "Input": ["84", "937"], // Each line of input as a separate element
    "Output": "85\n938" // The output expected
  },
  {
    "Input": ["24", "137"],
    "Output": "25\n138"
  },
  // .
  // .
  // .
  ]
}
```

Thats it! You're now done setting up a project for use! In the next section, you'll see how to use the cli for your testing!

## 📚 Usage 

Once you have set up your project, you're now ready to start testing! There are two ways in which you can test files -
```sh
$ addike-ct
```
Or
```sh
$ addike-ct -validate
```
The latter tests files _only after validating them_. If you use this, you will need to have a `validate.py` in your `config` directory. This is generally recommended when testing third party code.

You should see some output like this - 
<p align="center">
  <img src="./Pictures/output.png" />
</p>

If yes, congratulate yourself! You have successfully installed, and used `addike-ct`! You may also navigate to [localhost:8080](http://localhost:8080) to view your results in the form of a webpage, which was made with Svelte.
