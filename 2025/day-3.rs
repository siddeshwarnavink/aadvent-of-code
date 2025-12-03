use std::io::{self, BufRead};

fn max_number(digits: &[u8]) -> Vec<u8> {
    let mut drop = digits.len() - 12;
    let mut stack: Vec<u8> = Vec::with_capacity(12);

    for digit in digits.iter() {
        while drop > 0 && !stack.is_empty() && *stack.last().unwrap() < *digit {
            stack.pop();
            drop -= 1;
        }
        stack.push(*digit);
    }
    stack.truncate(12);
    stack
}

fn main() {
    let stdin = io::stdin();
    let mut ans: u64 = 0;

    for line in stdin.lock().lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        let digits: Vec<u8> = line.bytes().map(|b| b - b'0').collect();
        let max_digits = max_number(&digits);
        let mut max: u64 = 0;

        for digit in max_digits {
            max *= 10;
            max += u64::from(digit);
        }
        ans += max;
    }

    println!("{ans}");
}
