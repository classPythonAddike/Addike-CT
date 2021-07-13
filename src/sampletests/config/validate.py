import sys

file = sys.argv[1] # The file will be passed in as a command line argument

with open(file, "r") as f:
    code = f.read()

if "eval" in code:
    print("Program uses eval()!") # If the code does not satisfy the requirements, print out why
elif "exec" in code:
    print("Program uses exec()!")
elif "import" in code:
    print("Imports used!")
elif "getattr" in code:
    print("Program uses gettatr()")
elif "open" in code:
    print("Program accesses files!")
else:
    print("valid") # If the code is valid, print "valid" to the terminal, and nothing else
