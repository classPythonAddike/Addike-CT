# Program to generate random cases

import json, random

cases = {"Cases": []}

for i in range(20):
	case = {"Input": "", "Output": ""}

	for i in range(1):
	
		num = random.randint(1, 100)
		sol = num + 1 # Program should print the number inputted + 1
		case["Input"] += str(num) + "\n"
		case["Output"] += str(sol) + "\n"

	cases["Cases"] += [case]

with open("cases.json", "w") as f:
	json.dump(cases, f, indent=4)
