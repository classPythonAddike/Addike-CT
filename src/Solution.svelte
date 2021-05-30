<script>
  //import CircularProgress from '@smui/circular-progress';
  export let sol_name;
  export let time_taken;
  export let code_length;
  export let progress = 0;
  export let finished;
  export let passed;

  let progress_modified = 0;

  //  $: bar_color_value = (progress/100) * 125;
  $: bar_color_value =
    finished && passed
      ? "#079015"
      : finished && !passed
      ? "#B12000"
      : !finished && time_taken > 0
      ? "#074F90"
      : "#A5A5A5";

  $: progress_modified = !finished && time_taken == 0 ? 100 : progress; // gray bg if testing has not started
  $: progress_length = 503 - (progress_modified / 100) * 503;
</script>

<div class="card">
  <div class="name">
    <img src="personicon.svg" class="personicon" alt="icon" />
    <span>{sol_name}</span>
  </div>
  <div class="progress">
    <svg
      class="progressbar"
      stroke-dasharray="503"
      stroke-dashoffset={progress_length}
      stroke={bar_color_value}
    >
      <circle cx="90" cy="90" r="80" />
    </svg>
    <div class="progress">
      {progress}%
    </div>
  </div>
  <div class="info">
    <table class="information">
      <tr>
        <th class="first-item">Time Taken</th>
        <th class="last-item">Code Length</th>
      </tr>
      <tr>
        <td class="info-column">
          <div class="placeholder">
            <p class="data" style="display: inline;">{time_taken}</p>
            <span class="unit">secs</span>
          </div>
        </td>
        <td class="info-column-2">
          <div class="placeholder">
            <p class="data" style="display: inline;">{code_length}</p>
            <span class="unit">chars</span>
          </div>
        </td>
      </tr>
    </table>
  </div>
</div>

<style>
  .name {
    display: flex;
    justify-content: center;
    margin: 15px 0;
  }

  .personicon {
    height: 20px;
    margin-left: 5px;
  }

  span {
    padding-left: 5px;
  }

  .info-column-2 {
    padding-left: 3px;
    border: 0px;
  }

  table {
    border-collapse: collapse;
  }

  table td,
  table th {
    border: 0px solid black;
    border-right: 1px solid black;
  }
  table tr:first-child th {
    border-top: 0;
  }
  table tr:last-child td {
    border-bottom: 0;
  }
  table tr td:first-child,
  table tr th:first-child {
    border-right: 1px black solid;
  }

  .first-item {
    padding-right: 5px;
  }

  .last-item {
    border-right: 0px;
    padding-left: 5px;
  }

  .placeholder {
    width: 100%;
    display: flex;
    justify-content: center;
    font-family: "Montserrat", sans-serif;
    font-style: normal;
    font-weight: 600;
    font-size: 13.3px;
    line-height: 22px;
  }

  .card {
    height: 340px;
    width: 250px;
    background: rgba(220, 220, 220, 0.7);
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-direction: column;
    border-radius: 13px;
    font-family: "Montserrat", sans-serif;
    margin: 9px;
    box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.4);
  }

  .progress {
    margin: auto;
    width: 100%;
    height: 40%;
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: "Roboto", sans-serif;
    font-style: normal;
    font-weight: 500;
    font-size: 40.8px;
    line-height: 56px;
    text-align: center;
    color: #000000;
  }

  .progressbar {
    position: absolute;
    width: 180px;
    height: 180px;
    fill: transparent;
    stroke-width: 10px;
    stroke-linecap: round;
    transform: rotate(-90deg);
  }

  .info {
    height: 20%;
    width: 100%;
    display: flex;
    justify-content: space-around;
    align-items: center;
    margin: 10px 0;
  }

  .data {
    font-family: "Montserrat", sans-serif;
    font-style: normal;
    font-weight: 500;
    font-size: 15.3px;
    line-height: 21px;
    text-align: center;
  }

  .unit {
    color: #353535;
    font-family: "Montserrat", sans-serif;
    font-weight: 400;
  }
</style>
