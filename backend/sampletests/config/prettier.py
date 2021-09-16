import json, sys
import pandas as pd

with open(sys.argv[2], "r") as f:
    data = json.load(f)

for i in range(len(data)):
    del data[i]["Extension"]
    del data[i]["Language"]
    del data[i]["FinishedTesting"]
    del data[i]["Compiled"]

print(
    pd.DataFrame([list(i.values()) for i in data], columns=["File", "Completed Cases", "Passed", "Code Length", "Time"])
)
