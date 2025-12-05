use std::io;

fn main() {
    let stdin = io::stdin();

    let mut height: usize = 0;
    let mut width: usize = 0;

    let mut map: Vec<bool> = Vec::new();

    for line in stdin.lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        width = line.len();
        height += 1;

        for c in line.chars() {
            map.push(match c {
                '@' => true,
                _ => false,
            });
        }
    }

    let mut ans: usize = 0;

    let mut visited: Vec<(usize, usize)> = Vec::new();
    let mut checks: Vec<(usize, usize)> = Vec::new();

    loop {
        for (i, is_paper) in map.iter().enumerate() {
            if !*is_paper {
                continue;
            }

            let x = i / width;
            let y = i % width;

            checks.clear();

            if y + 1 < width {
                checks.push((x, y + 1));
            }
            if x + 1 < height {
                checks.push((x + 1, y));
                if y + 1 < width {
                    checks.push((x + 1, y + 1));
                }
            }
            if y > 0 {
                checks.push((x, y - 1));
                if x + 1 < height {
                    checks.push((x + 1, y - 1));
                }
            }
            if x > 0 {
                checks.push((x - 1, y));
                if y + 1 < width {
                    checks.push((x - 1, y + 1));
                }
            }
            if x > 0 && y > 0 {
                checks.push((x - 1, y - 1));
            }

            let mut count = 0;

            for (x, y) in &checks {
                let j = x * width + y;
                if let Some(is_paper) = map.get(j) {
                    if *is_paper {
                        count += 1;
                    }
                }
            }

            if count < 4 {
                ans += 1;
                visited.push((x, y));
            }
        }

        if visited.is_empty() {
            break;
        }

        while let Some((x, y)) = visited.pop() {
            let i = x * width + y;
            map[i] = false;
        }
    }

    println!("{ans}");
}
