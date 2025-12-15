use std::collections::HashMap;
use std::fmt;
use std::io;
use std::str::FromStr;

struct Point {
    x: i64,
    y: i64,
    z: i64,
}

impl Point {
    fn new(x: i64, y: i64, z: i64) -> Self {
        Self { x, y, z }
    }

    fn distance_with(&self, b: &Self) -> i64 {
        let dx = (b.x - self.x) as f64;
        let dy = (b.y - self.y) as f64;
        let dz = (b.z - self.z) as f64;
        let dist = dx * dx + dy * dy + dz * dz;
        dist.round() as i64
    }
}

impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{},{},{}", self.x, self.y, self.z)
    }
}

fn find_root(union_find: &mut HashMap<usize, usize>, i: usize) -> usize {
    if *union_find.get(&i).unwrap() == i {
        return i;
    }
    let root = find_root(union_find, *union_find.get(&i).unwrap());
    union_find.insert(i, root);
    root
}

fn union_sets(union_find: &mut HashMap<usize, usize>, a: usize, b: usize) -> bool {
    let root_a = find_root(union_find, a);
    let root_b = find_root(union_find, b);

    if root_a != root_b {
        union_find.insert(root_b, root_a);
        return true;
    }
    false
}

fn main() {
    let stdin = io::stdin();
    let mut points: Vec<Point> = Vec::new();

    for line in stdin.lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        let mut it = line.split(",").into_iter();
        let x = i64::from_str(it.next().unwrap_or("0")).unwrap();
        let y = i64::from_str(it.next().unwrap_or("0")).unwrap();
        let z = i64::from_str(it.next().unwrap_or("0")).unwrap();

        points.push(Point::new(x, y, z));
    }

    let mut distances: HashMap<(usize, usize), i64> = HashMap::new();

    for i in 0..points.len() {
        for j in 0..points.len() {
            if j == i {
                continue;
            }

            let a = usize::min(i, j);
            let b = usize::max(i, j);
            let key = (a, b);
            if !distances.contains_key(&key) {
                let value = points[a].distance_with(&points[b]);
                distances.insert(key, value);
            }
        }
    }

    let mut distances_vec = distances.iter().collect::<Vec<_>>();
    distances_vec.sort_by(|a, b| a.1.cmp(&b.1));

    let mut union_find: HashMap<usize, usize> = HashMap::new();
    for i in 0..points.len() {
        union_find.insert(i, i);
    }

    let mut count = 0;
    for item in distances_vec {
        if count >= 10 {
            break;
        }

        let i = item.0 .0;
        let j = item.0 .1;

        if union_sets(&mut union_find, i, j) {
            count += 1;
        }
    }

    let mut group_count: HashMap<usize, usize> = HashMap::new();

    for item in union_find.iter() {
        *group_count.entry(*item.1).or_insert(0) += 1;
    }

    let mut group_count: HashMap<usize, usize> = HashMap::new();
    for i in 0..points.len() {
        let final_root = find_root(&mut union_find, i);
        *group_count.entry(final_root).or_insert(0) += 1;
    }

    let mut group_count_vec = group_count.iter().collect::<Vec<_>>();
    group_count_vec.sort_by(|a, b| b.1.cmp(&a.1));

    for item in group_count_vec.iter() {
        println!("{item:?}");
    }
}
