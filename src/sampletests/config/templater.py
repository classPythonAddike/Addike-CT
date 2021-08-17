import templates

blacklisted = {
    ".c": ["include"],
    ".cpp": ["include"],
    ".cs": ["using"],
    ".go": ["import"],
    ".java": ["import"],
    ".js": ["require", "eval", "exec"],
    ".py": ["import", "eval", "exec", "getattr"],
    ".rb": ["require", "include", "eval", "exec"]
}

language_templates = {
    ".c": templates.c_template,
    ".cpp": templates.cpp_template,
    ".cs": templates.csharp_template,
    ".go": templates.go_template,
    ".py": templates.python_template,
    ".rb": templates.ruby_template
}

import os
from os import listdir
from os.path import isfile, join

files = [f for f in listdir(".") if isfile(join(".", f))]

for file in files:
    ext = os.path.splitext(file)[-1]

    