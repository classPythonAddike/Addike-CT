<script>
  let ChallengeName = "Simon Says";
  let ChallengeDescription =
    "Only what Simon said counts! On each test, you will be given a  string of any length. This string can be anything. Your job is to output what Simon said!";

  class Result {
    constructor(file, time, length, progress) {
      this.sol_name = file;
      this.time_taken = time;
      this.code_length = length;
      this.progress = progress;
    }
  }

  let result_array = [
    RandomResult(),
    RandomResult(),
    RandomResult(),
    RandomResult(),
    RandomResult(),
  ];

  import Solution from "./Solution.svelte";

  function RandomResult() {
    return new Result("123456789.py", 45.67, 183, 56.4);
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
  {#each result_array as { sol_name, time_taken, code_length, progress }, _}
    <Solution {sol_name} {time_taken} {code_length} {progress} />
  {/each}
</div>

<style>
  table {
    margin: 20px;
  }

  .container {
    display: flex;
    flex-wrap: wrap;
  }

  .heading {
    color: black;
    font-family: Roboto;
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
    font-family: Roboto;
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
    font-family: Roboto;
    font-style: normal;
    font-weight: normal;
    font-size: 27.2px;
    line-height: 37px;
    padding-right: 10px;
  }

  h5 {
    width: 100%;
    margin-left: 20px;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 500;
    font-size: 21.2px;
    line-height: 30px;

    color: #000000;
  }
</style>
