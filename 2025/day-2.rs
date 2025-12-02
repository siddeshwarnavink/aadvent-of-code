use std::io;
use std::str::FromStr;

fn made_of_substring(needle: &str, hay: &str) -> bool {
    let n = needle.len();
    let mut i = 0;
    while i < hay.len() {
        if let Some(win) = hay.get(i..i+n) {
            if needle != win {
                return false;
            }
            i += needle.len()
        } else {
            return false;
        }
    }
    true
}

#[test]
fn made_of_substring_test_positive() {
    assert!(made_of_substring("a", "aaa"));
    assert!(made_of_substring("abc", "abcabc"));
    assert!(made_of_substring("abc", "abcabcabc"));
}

#[test]
fn made_of_substring_test_negative() {
    assert!(!made_of_substring("abc", "ab"));
    assert!(!made_of_substring("abc", "abcababc"));
    assert!(!made_of_substring("abc", "abcabcab"));
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut buf = String::new();

    stdin.read_line(&mut buf)?;
    let _ = buf.pop();

    let mut ans: i64 = 0;

    for pair in buf.split(",") {
        let start = i64::from_str(pair.split("-").nth(0).unwrap_or("0"))
            .unwrap_or(0);
        let end = i64::from_str(pair.split("-").nth(1).unwrap_or("0"))
            .unwrap_or(0);

        for i in start..(end + 1) {
            let s = i.to_string();
            let n = s.len();

            if n % 2 == 0 {
                let a = &s[..n/2];
                let b = &s[n/2..];
                if a == b {
                    ans += i;
                    continue;
                }
            }

            for len in 1..(n / 2)+1 {
                let win = &s.as_str()[..len];
                let rest = &s.as_str()[len..];

                if made_of_substring(win, rest) {
                    ans += i;
                    break;
                }
            }
        }
    }
    println!("{ans}");

    Ok(())
}
