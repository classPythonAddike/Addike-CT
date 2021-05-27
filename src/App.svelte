<script>
  import Solution from "./Solution.svelte";

  let ChallengeName = "---"; // "Simon Says"
  let ChallengeDescription = "---"
    // "Only what Simon said counts! On each test, you will be given a  string of any length. This string can be anything. Your job is to output what Simon said!";

  class Result {
    constructor(file, progress, time, length) {
      this.sol_name = file;
      this.time_taken = time;
      this.code_length = length;
      this.progress = progress;
    }
  }

  let result_array = []

  async function getStartup() {
    try {
      const response = await fetch('http://127.0.0.1:5000/startup')
      const data = await response.json()

      ChallengeName = data.name
      ChallengeDescription = data.description

      document.title = `${data.name} | Challenge Tester`

      result_array = data.files.map((file) => {
        return new Result(file.name, file.progress, file.time, file.length)
      })
    }
    catch {
      console.log("Unable to connect to backend.")
    }
  }

  getStartup()

  const interval = setInterval(() => {getUpdate()}, 1000)
  async function getUpdate() {
    try {
      const response = await fetch('http://127.0.0.1:5000/progress')
      const data = await response.json()

      data.values.forEach((progress, index) => {
        result_array[index].progress = progress
      })
      data.times.forEach((time, index) => {
        result_array[index].time = time
      })
    }
    catch {
      console.log("Unable to connect to backend.")
      clearInterval(interval)
    }
  }
</script>

<table class="heading">
  <tr>
    <td class="title">
      <h1>{ChallengeName}</h1>
      <h4>Weekly Challenge</h4>
    </td>
    <td class="desc"><h5>{ChallengeDescription}</h5></td>
  </tr>
</table>

<div class="container">
  {#each result_array as sol}
    <Solution {...sol} />
  {/each}
</div>

<style>
  table {
    margin: 20px;
  }

  .container {
    display: flex;
    flex-wrap: wrap;
    margin: 0 auto;
  }

  .heading {
    color: black;
    font-family: 'Roboto', sans-serif;
  }

  .title {
    margin-left: 20px;
    border-right: 1px black solid;
    width: max-content;
  }

  td :last-child {
    margin-left: 0px;
  }

  .desc {
    padding-left: 10px;
  }

  h1 {
    width: max-content;
    margin-right: 10px;
    font-family: 'Roboto', sans-serif;
    font-style: normal;
    font-weight: normal;
    font-size: 45px;
    line-height: 60px;
    color: #000000;
  }

  h4 {
    width: max-content;
    text-align: left;
    margin-left: 0;
    font-family: 'Roboto', sans-serif;
    font-style: normal;
    font-weight: normal;
    font-size: 27.2px;
    line-height: 37px;
    padding-right: 10px;
  }

  h5 {
    width: 100%;
    margin-left: 20px;
    font-family: 'Montserrat', sans-serif;
    font-style: normal;
    font-weight: 500;
    font-size: 21.2px;
    line-height: 30px;

    color: #000000;
  }
</style>
