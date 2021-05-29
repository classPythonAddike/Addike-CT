<script>
  import Solution from "./Solution.svelte";

  let ChallengeName = "---";
  let ChallengeDescription = "---";

  class Result {
    constructor(file, progress, time, length) {
      this.sol_name = file;
      this.time_taken = time;
      this.code_length = length;
      this.progress = progress;
    }
  }

  let result_array = [];

  async function getStartup() {
    try {
      const response = await fetch("http://localhost:8080/startup");
      const data = await response.json();

      ChallengeName = data.Name;
      ChallengeDescription = data.Description;

      document.title = `${data.Name} | Challenge Tester`;

      result_array = data.Files.map((file) => {
        return new Result(file.File, 0, 0, file.CodeLength);
      });
    } catch {
      console.log("Unable to connect to backend.");
    }
  }

  getStartup();

  const interval = setInterval(() => {
    getUpdate();
  }, 50);
  async function getUpdate() {
    try {
      const response = await fetch("http://localhost:8080/progress");
      const data = await response.json();

      data.Values.forEach((progress, index) => {
        result_array[index].progress = progress.toFixed(1);
      });
      data.Times.forEach((time, index) => {
        result_array[index].time_taken = time.toFixed(1);
      });
    } catch {
      console.log("Unable to connect to backend.");
      clearInterval(interval);
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
  {#each result_array as soln}
    <Solution {...soln} />
  {/each}
</div>

<style>
  table {
    margin: 20px;
  }

  .container {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    margin: 0 auto;
  }

  .heading {
    color: black;
    font-family: "Roboto", sans-serif;
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
    font-family: "Roboto", sans-serif;
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
    font-family: "Roboto", sans-serif;
    font-style: normal;
    font-weight: normal;
    font-size: 27.2px;
    line-height: 37px;
    padding-right: 10px;
  }

  h5 {
    width: 100%;
    margin-left: 20px;
    font-family: "Montserrat", sans-serif;
    font-style: normal;
    font-weight: 500;
    font-size: 21.2px;
    line-height: 30px;

    color: #000000;
  }
</style>
