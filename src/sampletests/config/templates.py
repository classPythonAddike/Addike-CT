c_template = """
#include <stdio.h>

// Code goes here
[code]

int main() {
    printf(solution(gets()));
}
"""

cpp_template = """
#include <iostream>
#include <string>
#include <vector>

// Code goes here
[code]

int main() {
    var input string;
    std::cin >> input;
    std::cout << solution();
}
"""

csharp_template = """
using System;

class Program {    

    static void Main(string[] args)
    {
        var input string = Console.ReadLine();
        Console.Write(solution(input));
    }

    // Code goes here
    [code]
}
"""

go_template = """
package main

import "fmt"

func main() {
    var input string
    fmt.Scan(&first)
    fmt.Print(solution(input))
}

// Code goes here
[code]
"""

python_template = """
# Code goes here
[code]

inp = input()
print(solution(inp), end="")
"""

ruby_template = """
# Code goes here
[code]

input = gets
prints solution(input)
"""