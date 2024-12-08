use std::{fs, path::Path};

fn parse_data(data_string: &String) -> Vec<Vec<i32>> {
    data_string
        .lines()
        .fold(Vec::new(), |mut acc, line| {
            acc.push(line
                .split_whitespace()
                .fold(Vec::new(), |mut acc, value| {
                    acc.push(value.parse().unwrap());
                    acc
                }));
            acc
        })
}

fn is_safe(report: &[i32]) -> bool {
    let mut n = report[0];
    let mut is_ascending: bool = false;
    
    if n < report[1] {
        is_ascending = true;
    }

    for i in 1..report.len() {
        let k = report[i];
        let abs_diff = (n - k).abs();
        let in_range = abs_diff < 1 || 3 < abs_diff;

        if is_ascending && (n >= k || in_range) {
            return false;
        }
        else if !is_ascending && (n < k || in_range) {
            return false;
        }
        
        n = report[i];
    }
       
    true
}

pub fn day2_solve() {
    let data_string = fs::read_to_string(Path::new("data/day2_main.txt")).unwrap();
    let parsed_data = parse_data(&data_string);
    let safe_num = parsed_data
        .iter()
        .fold(0, |mut acc, report| {
            if is_safe(&report) {
                acc += 1;
            }
            acc
        });
    println!("{}", safe_num);
}
