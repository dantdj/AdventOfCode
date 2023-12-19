use std::{
    fs::File,
    io::{BufRead, BufReader},
    time::Instant,
};

fn main() {
    measure_time(part_one);
    measure_time(part_two);
}

fn part_one() -> String {
    let input = read_input("real-input");

    get_trees_on_slope(&input, 1, 3).to_string()
}

fn part_two() -> String {
    let input = read_input("real-input");

    let slope1_trees = get_trees_on_slope(&input, 1, 1);
    let slope2_trees = get_trees_on_slope(&input, 1, 3);
    let slope3_trees = get_trees_on_slope(&input, 1, 5);
    let slope4_trees = get_trees_on_slope(&input, 1, 7);
    let slope5_trees = get_trees_on_slope(&input, 2, 1);

    (slope1_trees * slope2_trees * slope3_trees * slope4_trees * slope5_trees).to_string()
}

/// Given a 2D Vector and the required down and right step values, calculates how many trees
/// will be collided with when traversing the resulting slope
fn get_trees_on_slope(input: &Vec<Vec<char>>, down_step: usize, right_step: usize) -> i64 {
    let mut right_ctr = 0;
    let mut down_ctr = 0;
    let mut tree_counter = 0;

    while down_ctr < input.len() - 1 {
        if down_ctr + down_step >= input.len() {
            break;
        }

        down_ctr += down_step;

        // If the right counter has reached the end of the array,
        // loop back around
        right_ctr = if right_ctr + right_step >= input[0].len()  {
            // Step value minus the difference between the end of the array
            // and the current right counter value
            right_step - ((input[0].len()) - right_ctr)
        } else {
            right_ctr + right_step
        };

        if input[down_ctr][right_ctr] == '#' {
            tree_counter += 1
        }
    }

    tree_counter
}

fn read_input(file_path: &str) -> Vec<Vec<char>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);

    // Split the input string into lines
    let lines: Vec<_> = reader
        .lines()
        .collect::<Result<_, _>>()
        .expect("Failed to read lines");

    // Determine the dimensions of the 2D array
    let num_lines = lines.len();
    let max_line_length = lines.iter().map(|line| line.len()).max().unwrap_or(0);

    // Create a 2D array of characters
    let mut char_array: Vec<Vec<char>> = vec![vec![' '; max_line_length]; num_lines];

    // Populate the 2D array with characters
    for (i, line) in lines.iter().enumerate() {
        for (j, c) in line.chars().enumerate() {
            char_array[i][j] = c;
        }
    }

    char_array
}

fn measure_time(f: fn() -> String) {
    let start = Instant::now();
    let result: String = f();
    let duration = start.elapsed();

    println!("{} - {}µs", result, duration.as_micros())
}
