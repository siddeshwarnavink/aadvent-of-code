use std::io::{self, BufRead};
use std::str::FromStr;

fn main() {
    let stdin = io::stdin();
    let mut point: i32 = 50;
    let mut ans: i32 = 0;

    for line in stdin.lock().lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        let direction = &line[0..1];
        let count = i32::from_str(&line[1..]).unwrap_or(0);

        for _ in 0..count {
            point = match direction {
                "L" => ((point - 1) + 100) % 100,
                "R" => ((point + 1) + 100) % 100,
                _ => point,
            };
            if point == 0 {
                ans += 1;
            }
        }
    }

    println!("{ans}");
}
