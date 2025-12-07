use std::io;
use std::str::FromStr;

fn main() {
    let stdin = io::stdin();
    let mut data: Vec<Vec<String>> = Vec::new();
    let mut ans = 0i64;

    for line in stdin.lines() {
        let line = match line {
            Ok(l) => l,
            Err(_) => continue,
        };

        let mut i = 0;

        for part in line.split(" ") {
            if part.len() < 1 {
                continue;
            }

            if let Ok(_) = i64::from_str(part) {
                if let Some(bucket) = data.get_mut(i) {
                    bucket.push(part.to_string());
                } else {
                    data.insert(i, vec![part.to_string()]);
                }
            } else {
                let digit_count = {
                    let mut max = 0usize;
                    for num in &data[i] {
                        max = usize::max(max, num.len());
                    }
                    max
                };

                match part {
                    "+" => {
                        let mut sum = 0i64;

                        for j in (0..digit_count).rev() {
                            let mut num_buf = String::new();

                            for num in &data[i] {
                                if let Some(digit) = num.get(j..j + 1) {
                                    num_buf.push_str(digit);
                                }
                            }

                            let num = i64::from_str(num_buf.as_str()).unwrap();
                            sum += num;
                        }

                        ans += sum;
                    }
                    "*" => {
                        let mut prod = 1i64;

                        for j in (0..digit_count).rev() {
                            let mut num_buf = String::new();

                            for num in &data[i] {
                                let mut num_with_padding = num.to_string();
                                while num_with_padding.len() < digit_count {
                                    num_with_padding.insert(0, 'a');
                                }

                                if let Some(digit) = num_with_padding.get(j..j + 1) {
                                    if digit != "a" {
                                        num_buf.push_str(digit);
                                    }
                                }
                            }

                            let num = i64::from_str(num_buf.as_str()).unwrap();
                            prod *= num;
                        }

                        ans += prod;
                    }
                    _ => unreachable!(),
                }
            }

            i += 1;
        }
    }

    println!("{ans}");
}
