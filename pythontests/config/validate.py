import sys

file = sys.argv[1]

with open(file, "r") as f:
    code = f.read()

if "eval" in code and file != "mysolution.py":
    print("Program uses eval()!")
elif "exec" in code:
    print("Program uses exec()!")
elif "import" in code:
    print("Imports used!")
elif "getattr" in code:
    print("Program uses gettatr()")
elif "open" in code:
    print("Program accesses files!")
else:
    print("valid")
