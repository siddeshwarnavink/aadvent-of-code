use std::collections::{HashSet, HashMap};
use std::io;

fn main() {
    let stdin = io::stdin();
    let mut start: Option<(usize, usize)> = None;
    let mut splitters: HashSet<(usize, usize)> = HashSet::new();
    let mut height = 0usize;

    for (i, line) in stdin.lines().enumerate() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };
        if line.len() < 1 { break; }

        for (j, c) in line.chars().enumerate() {
            match c {
                '.' => continue,
                'S' => {
                    assert_eq!(start, None);
                    start = Some((i, j));
                }
                '^' => {
                    splitters.insert((i, j));
                }
                _ => unreachable!(),
            }
        }

        height += 1;
    }

    let start = start.unwrap();

    let mut beams: HashMap<(usize, usize), usize> = HashMap::new();
    beams.insert(start, 1);

    for _ in 0..height {
        let mut next_beams: HashMap<(usize, usize), usize> = HashMap::new();

        for (beam, count) in beams.drain() {
            let x = beam.0;
            let y = beam.1;

            if splitters.contains(&(x, y)) {
                *next_beams.entry((x + 1, y - 1)).or_insert(0) += count;
                *next_beams.entry((x + 1, y + 1)).or_insert(0) += count;
            } else {
                *next_beams.entry((x + 1, y)).or_insert(0) += count;
            }
        }

        beams = next_beams;
    }

    let total_paths: usize = beams.values().sum();
    println!("{}", total_paths);
}
