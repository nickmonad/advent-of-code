import gleam/dict.{type Dict}
import gleam/int
import gleam/list
import gleam/string
import gleam/result
import gleam/option.{Some, None}

pub fn part1(input: List(String)) -> String {
    let #(left, right) = input
        |> list.map(fn(line) { string.split(line, on: " ") })
        |> list.filter(fn(line) { !list.is_empty(line) })
        |> list.filter_map(fn(pair) {
            // take each 2-element list, map to 2-element tuple of ints,
            // and "unzip" to get a left list and right list
            case pair {
                [l, r] -> {
                    // assume valid input, can always convert to int
                    let assert Ok(l) = int.base_parse(l, 10)
                    let assert Ok(r) = int.base_parse(r, 10)

                    Ok(#(l, r))
                }
                _ -> Error("expected sets of 2")
            }
        })
        |> list.unzip

    let left = left |> list.sort(by: int.compare)
    let right = right |> list.sort(by: int.compare)

    let sum = list.zip(left, right)
        |> list.map(fn(x) { int.absolute_value(x.0 - x.1) })
        |> int.sum

    sum |> int.to_string
}

pub fn part2(input: List(String)) -> String {
    let #(left, right) = input
        |> list.map(fn(line) { string.split(line, on: " ") })
        |> list.filter(fn(line) { !list.is_empty(line) })
        |> list.filter_map(fn(pair) {
            // take each 2-element list, map to 2-element tuple of ints,
            // and "unzip" to get a left list and right list
            case pair {
                [l, r] -> {
                    // assume valid input, can always convert to int
                    let assert Ok(l) = int.base_parse(l, 10)
                    let assert Ok(r) = int.base_parse(r, 10)

                    Ok(#(l, r))
                }
                _ -> Error("expected sets of 2")
            }
        })
        |> list.unzip

    // build a dict where key = number and value = count (of numbers in right list)
    let counts: Dict(Int, Int) = right
        |> list.fold(dict.new(), fn(counts, r) {
            dict.upsert(counts, r, fn(current) {
                case current {
                    Some(val) -> val + 1
                    None -> 1
                }
            })
        })

    // sum every item in the left,
    // multiplied by its occurances in the right
    let sum: Int = left |> list.fold(0, fn(acc, n) {
        acc + { n * { dict.get(counts, n) |> result.unwrap(0) } }
    })

    sum |> int.to_string
}
