use regex::Regex;
use std::{
    fs::File,
    io::{self, BufRead, BufReader},
    time::Instant, clone,
};

fn main() {
    measure_time(part_one);
    measure_time(part_two);
}

fn part_one() -> String {
    let input = p1_read_input("real-input").unwrap();
    let mut valid_count = 0;
    for password_details in input.iter() {
        let mut count = 0;
        for character in password_details.password.chars() {
            if character == password_details.required_character {
                count += 1;
            }
        }

        // This might be overly fancy, but essentially we construct an inclusive range,
        // and check if the count of characters is within that range.
        if (password_details.minimum_appearances..=password_details.maximum_appearances)
            .contains(&count)
        {
            valid_count += 1;
        }
    }

    valid_count.to_string()
}

fn part_two() -> String {
    let input = p2_read_input("real-input").unwrap();
    let mut valid_count = 0;

    for details in input.iter() {
        let mut first_matches = false;
        let mut second_matches = false;

        if let Some(ch) = details.password.chars().nth(details.first_position as usize - 1) {
            first_matches = ch == details.required_character;
        }

        if let Some(ch) = details.password.chars().nth(details.second_position as usize - 1) {
            second_matches = ch == details.required_character;
        }

        // This is basically an exclusive OR - if both are false, it's not valid, and if both are true, it's not valid
        if first_matches != second_matches {
            valid_count += 1;
        }
    }

    valid_count.to_string()
}

fn measure_time(f: fn() -> String) {
    let start = Instant::now();
    let result: String = f();
    let duration = start.elapsed();

    println!("{} - {}µs", result, duration.as_micros())
}

fn p1_read_input(file_path: &str) -> io::Result<Vec<PartOnePasswordInput>> {
    let file = File::open(file_path)?;
    let reader = BufReader::new(file);

    let re = Regex::new(r"(\d+)-(\d+) (\w): (\w+)").map_err(|e| {
        io::Error::new(
            io::ErrorKind::InvalidData,
            format!("Regex compilation error: {}", e),
        )
    })?;

    let password_inputs: Result<Vec<_>, io::Error> =
        reader.lines().map(|line| p1_parse_line(re.clone(), &line?)).collect();

    password_inputs
}

fn p1_parse_line(re: Regex, input: &str) -> Result<PartOnePasswordInput, io::Error> {
    let captures = re
        .captures(input)
        .ok_or_else(|| io::Error::new(io::ErrorKind::InvalidData, "Regex match failed"))?;

    let minimum = captures[1]
        .parse()
        .map_err(|e| io::Error::new(io::ErrorKind::InvalidData, e))?;
    let maximum = captures[2]
        .parse()
        .map_err(|e| io::Error::new(io::ErrorKind::InvalidData, e))?;
    let required_char = captures[3]
        .chars()
        .next()
        .ok_or_else(|| io::Error::new(io::ErrorKind::InvalidData, "Missing required character"))?;
    let password = captures[4].to_string();

    Ok(PartOnePasswordInput {
        minimum_appearances: minimum,
        maximum_appearances: maximum,
        required_character: required_char,
        password,
    })
}

struct PartOnePasswordInput {
    password: String,
    required_character: char,
    minimum_appearances: i32,
    maximum_appearances: i32,
}


fn p2_read_input(file_path: &str) -> io::Result<Vec<PartTwoPasswordInput>> {
    let file = File::open(file_path)?;
    let reader = BufReader::new(file);

    let re = Regex::new(r"(\d+)-(\d+) (\w): (\w+)").map_err(|e| {
        io::Error::new(
            io::ErrorKind::InvalidData,
            format!("Regex compilation error: {}", e),
        )
    })?;

    let password_inputs: Result<Vec<_>, io::Error> =
        reader.lines().map(|line| p2_parse_line(re.clone(), &line?)).collect();

    password_inputs
}

fn p2_parse_line(re: Regex, input: &str) -> Result<PartTwoPasswordInput, io::Error> {
    let captures = re
        .captures(input)
        .ok_or_else(|| io::Error::new(io::ErrorKind::InvalidData, "Regex match failed"))?;

    let first_position = captures[1]
        .parse()
        .map_err(|e| io::Error::new(io::ErrorKind::InvalidData, e))?;
    let second_position = captures[2]
        .parse()
        .map_err(|e| io::Error::new(io::ErrorKind::InvalidData, e))?;
    let required_char = captures[3]
        .chars()
        .next()
        .ok_or_else(|| io::Error::new(io::ErrorKind::InvalidData, "Missing required character"))?;
    let password = captures[4].to_string();

    Ok(PartTwoPasswordInput {
        first_position,
        second_position,
        required_character: required_char,
        password,
    })
}

struct PartTwoPasswordInput {
    password: String,
    required_character: char,
    first_position: i32,
    second_position: i32,
}
