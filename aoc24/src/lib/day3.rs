use std::{fs, path::Path};

fn clean_chunk(chunk: &str) -> String {
    chunk.chars().take_while(|c| *c != ')').collect()
}

fn parse_data(data_string: &str) -> Vec<(i32, i32)> {
    let split_by_mul: Vec<&str> = data_string.split("mul").collect();

    split_by_mul
        .iter()
        .fold(Vec::new(), |mut acc, chunk| {
            let cleaned_chunk = clean_chunk(&chunk[1..chunk.len()]);
            let split_by_comma: Vec<&str> = cleaned_chunk.split(",").collect();

            if split_by_comma.len() == 2 {
                let x: i32 = split_by_comma[0].parse().unwrap_or(0);
                let y: i32 = split_by_comma[1].parse().unwrap_or(0);
                acc.push((x, y));
            }

            acc
        })
}

pub fn day3_solve() {
    let data_string = fs::read_to_string(Path::new("data/day3_main.txt")).unwrap();
    let instructions = parse_data(&data_string);
    let result = instructions
        .iter()
        .fold(0, |mut acc, inst| {
            acc += inst.0 * inst.1;
            acc
        });
    println!("{}", result);
}
