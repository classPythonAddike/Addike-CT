<script>
  import Solution from "./Solution.svelte";

  let ChallengeName = "---";
  let ChallengeDescription = "----------";

  let error_happened = false;

  class Result {
    constructor(file, progress, time, length, finished, passed) {
      this.sol_name = file;
      this.time_taken = time;
      this.code_length = length;
      this.progress = progress;
      this.finished = finished;
      this.passed = passed;
    }
  }

  async function fetchSpecial(resource, options) {
    const { timeout } = options;

    const controller = new AbortController();
    const id = setTimeout(() => {
      controller.abort();
    }, timeout);

    const response = await fetch(resource, {
      ...options,
      signal: controller.signal,
    });
    clearTimeout(id);

    return response;
  }



  let result_array = [];

  async function getStartup() {
    try {
      const response = await fetchSpecial("http://localhost:8080/startup", {
        timeout: 120,
      });
      const data = await response.json();

      ChallengeName = data.Name;
      ChallengeDescription = data.Description;

      document.title = `${data.Name} | Challenge Tester`;

      result_array = data.Files.map((file) => {
        return new Result(file.File, 0, 0, file.CodeLength, false, false);
      });
    } catch(err) {
      handleError(err);
    }
  }

  getStartup();

  let loop = undefined;

  setTimeout(() => {
    if (!error_happened) {
      loop = setInterval(() => {
        getUpdate();
      }, 50);
    }
  }, 150);

  async function getUpdate() {
    try {
      const response = await fetchSpecial("http://localhost:8080/progress", {
        timeout: 50,
      });
      const data = await response.json();

      data.Values.forEach((progress, index) => {
        result_array[index].progress = progress.toFixed(1);
      });
      data.Times.forEach((time, index) => {
        result_array[index].time_taken = time.toFixed(1);
      });
      data.FinishedTesting.forEach((finished, index) => {
        result_array[index].finished = finished;
      });
      data.Passed.forEach((pass, index) => {
        result_array[index].passed = pass;
      });
    } catch (err) {
      handleError(err);
    }
  }

  function handleError(error) {
    console.log(loop);
    if (!error_happened) error_happened = !error_happened;
    if (loop) clearInterval(loop);
    if (error.name == "AbortError")
      alert(
        "Backend could be offline.\nCheck if it's still running and reload the page"
      );
    else {
      alert(
        `Something has gone wrong.\nTry checking if the backend is running or reload this page.\nError: ${error}`
      );
      console.log(`An error has occurred. Error: ${error}`);
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
