use std::{fs, path::Path};

fn get_left_right_sorted_values(data_str: &str) -> (Vec<i32>, Vec<i32>) {
    let mut left_vals: Vec<i32> = Vec::new();
    let mut right_vals: Vec<i32> = Vec::new();

    for (idx, slice) in data_str.split_whitespace().enumerate() {
        if idx % 2 == 0 {
            left_vals.push(slice.parse().unwrap());
        }
        else {
            right_vals.push(slice.parse().unwrap());
        }
    } 
    
    left_vals.sort();
    right_vals.sort();

    (left_vals, right_vals)
}

fn calculate_distances(lefts: &Vec<i32>, rights: &Vec<i32>) -> Vec<i32> {
    (0..lefts.len())
        .fold(Vec::new(), |mut acc, idx| {
            acc.push((lefts[idx] - rights[idx]).abs());
            acc
        })
}

pub fn day1_solve() {
    let data_string = fs::read_to_string(Path::new("data/day1_main.txt")).unwrap();     
    let left_right_vals = get_left_right_sorted_values(&data_string);
    let lefts = left_right_vals.0;
    let rights = left_right_vals.1;
    let distances = calculate_distances(&lefts, &rights);
    println!("{:?}", distances.iter().sum::<i32>());
}
