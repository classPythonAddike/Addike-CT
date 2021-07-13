import json, random

cases = {"Cases": []}

for i in range(300):
	case = {"Input": "4\n", "Output": ""}
	for i in range(4):
		num = random.randint(1, 1000)
		sol = num * (num - 1)//2
		case["Input"] += str(num) + "\n"
		case["Output"] += str(sol) + "\n"

	cases["Cases"] += [case]

with open("cases.json", "w") as f:
	json.dump(cases, f, indent=4)
