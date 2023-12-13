use std::{
    fs::File,
    io::{self, BufRead, BufReader},
    time::Instant,
};

fn main() {
    measure_time(part_one);
    measure_time(part_two);
}

fn part_one() -> String {
    let input = read_input("real-input").unwrap();

    let mut result = 0;
    
    // For each value, go through each other line and add them together to see
    // if they're the right value. We never have to go back, as we would've already tried
    // those combinations
    // We use the indexes as a little optimization as a result
    for (i, line) in input.iter().enumerate() {
        for j in i..input.len() {
            if line + input[j] == 2020 {
                result = line * input[j];
            }
        }
    }
    
    result.to_string()
}

fn part_two() -> String {
    let input = read_input("real-input").unwrap();

    let mut result = 0;

    for (i, line) in input.iter().enumerate() {
        for j in i..input.len() {
            for k in j..input.len() {
                if line + input[j] + input[k] == 2020 {
                    result = line * input[j] * input[k];
                }
            }
           
        }
    }

    result.to_string()
}

fn read_input(file_path: &str) -> io::Result<Vec<i32>> {
    let file = File::open(file_path)?;
    let reader = BufReader::new(file);

    let numbers: Result<Vec<i32>, _> = reader
        .lines()
        .map(|line| {
            line.and_then(|s| {
                s.trim()
                    .parse()
                    .map_err(|e| io::Error::new(io::ErrorKind::InvalidData, e))
            })
        })
        .collect();

    numbers
}

fn measure_time(f: fn() -> String) {
    let start = Instant::now();
    let result: String = f();
    let duration = start.elapsed();

    println!("{} - {}µs", result, duration.as_micros())
}
