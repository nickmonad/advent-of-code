import gleam/result
import gleam/int.{to_string}
import argv
import gleam/string
import gleam/list
import gleam/io
import simplifile as file
import glint

import solutions/day1

pub type Solution {
    Solution(day: Int, part: Int, do: fn(List(String)) -> String)
}

pub fn main() {
    glint.new()
        |> glint.with_name("aoc")
        |> glint.pretty_help(glint.default_pretty_help())
        |> glint.add(at: [], do: run())
        |> glint.run(argv.load().arguments)
}

fn run() -> glint.Command(Nil) {
    use day <- glint.named_arg("day")
    use part <- glint.flag(glint.int_flag("part") |> glint.flag_default(1))
    use named, _, flags <- glint.command()

    let assert Ok(day) = day(named) |> int.base_parse(10)
    let assert Ok(part) = part(flags)

    let solutions = [
        Solution(day: 1, part: 1, do: day1.part1),
        Solution(day: 1, part: 2, do: day1.part2),
    ]

    let path = "input/day" <> to_string(day)
    let assert Ok(input) = file.read(from: path) |> result.map(string.split(_, on: "\n"))

    let solution = list.find(solutions, fn(s) { s.day == day && s.part == part })
    let result = case solution {
        Ok(s) -> s.do(input)
        Error(_) -> {
            "error: no such solution, day " <> to_string(day) <> " part " <> to_string(part)
       }
    }

    io.println(result)
}
