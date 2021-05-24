<script>
    import CircularProgress from '@smui/circular-progress';
    export let sol_name = '123456789.py'
    export let time_taken = 5.6
    export let code_length = 55
    export let progress = 50
</script>
  
<div class="card">
    <div class="name">
        Solution: {sol_name}
    </div>
    <div class="progress">
        {progress}%
    </div>
    <div class="info">
        <div class="time"> 
            Time Taken
            <br/>
            {time_taken}s
        </div>
        <div class="code">
            Code Length
            <br/>
            {code_length} chars
        </div>
    </div>
</div>

<style>
    .card {
        height: 300px;
        width: 300px;
        background: #fff;
        display: flex;
        justify-content: space-around;
        align-items: center;
        flex-direction: column;
        border-radius: 10px;
        font-family: 'Montserrat', sans-serif;
        margin: 20px;
    }
    .progress {
        width: 100%;
        height: 40%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .info {
        height: 20%;
        width: 100%;
        display: flex;
        justify-content: space-around;
        align-items: center;
    }
    .code {
        text-align: center;
        display: flex;
        justify-content: center;
        align-items: center;
        border-left: 1px solid black;
    }
    .time {
        text-align: center;
        display: flex;
        justify-content: center;
        align-items: center;
        border-right: 1px solid black;
    }
    .code, .time {
        width: 50%;
        height: 100%;
    }
</style>