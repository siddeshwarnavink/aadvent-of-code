use std::io;
use std::str::FromStr;

fn fold_ranges(unsorted_ranges: &[(u64, u64)]) -> Vec<(u64, u64)> {
    let mut ranges = Vec::from(unsorted_ranges);
    ranges.sort();

    let mut res: Vec<(u64, u64)> = Vec::new();
    let mut i = 0;

    while i < ranges.len() {
        let start = ranges[i].0;
        let mut end = ranges[i].1;

        if !res.is_empty() && res[res.len() - 1].1 >= end {
            i += 1;
            continue;
        }

        let mut j = i + 1;
        while j < ranges.len() {
            if ranges[j].0 <= end {
                end = u64::max(end, ranges[j].1);
            }
            j += 1;
        }
        res.push((start, end));
        i += 1;
    }

    res
}

#[test]
fn fold_ranges_test() {
    let ranges = [(7, 8), (1, 5), (2, 4), (4, 6)];
    let new_ranges = fold_ranges(&ranges);

    assert_eq!(new_ranges.len(), 2);
    assert_eq!(new_ranges[0].0, 1);
    assert_eq!(new_ranges[0].1, 6);
    assert_eq!(new_ranges[1].0, 7);
    assert_eq!(new_ranges[1].1, 8);
}

fn main() {
    let stdin = io::stdin();
    let mut ranges: Vec<(u64, u64)> = Vec::new();

    for line in stdin.lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        if line.len() < 1 {
            break;
        }

        let mut it = line.split("-");
        let a = u64::from_str(it.next().unwrap_or("0")).unwrap();
        let b = u64::from_str(it.next().unwrap_or("0")).unwrap();

        ranges.push((a, b));
    }

    let new_ranges = fold_ranges(&ranges);
    drop(ranges);

    let mut ans: u64 = 0;
    for range in new_ranges {
        ans += range.1 - range.0 + 1;
    }
    println!("{}", ans);
}
