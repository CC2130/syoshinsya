extern crate rand;
use rand::{thread_rng, Rng};

fn main() {
    println!("Hello, world!");
    demo();
}

fn demo() {
    let mut v1 = crate_range_vector();
    let mut v2 = crate_range_vector();
    println!("v1 = {:?}\nv2 = {:?}", v1, v2);

    let v = merge(&mut v1, &mut v2);
    println!("merged v is: {:?}", v);
}

// 将有序的 v2 插入有序的 v1 里，并保持有序状态
fn merge(v1: &mut Vec<i32>, v2: &mut Vec<i32>) -> Vec<i32> {
    let mut v: Vec<i32> = Vec::new();
    let mut start = 0;          // v2 的起始位置
    let mut end: usize;         // v2 的结束位置
    let mut _insert: usize;     // 由此，插入 v2 数据 v2[start..end]

    // 以 v1 填充 v
    for _i in v1 {
        v.push(*_i);
    }
    if v2.is_empty() {
        return v
    }

    loop {
        let mut _insert = insert_position(&v, v2[start]);
        if _insert == v.len() {
            end = v2.len();
        } else {
            end = insert_position(v2, v[_insert]);
        }
        
        for i in &v2[start..end] {
            v.insert(_insert, *i);
            _insert += 1;
        }
        if end == v2.len() {
            break
        }
        start = end;
    }

    v
}


// 若把 n 插入 v 中, 返回位置，使得 v[p] <= n < v[p+1]
// 二分法查找
#[allow(dead_code)]
fn insert_position(v: &Vec<i32>, n: i32) -> usize {
    let len = v.len();
    // v 可能为空
    if len == 0 {
        return 0
    }
    let mut max = len - 1;
    let mut min = 0;
    let mut p;

    loop {
        p = (max + min) / 2;
        if n < v[p] {
            max = p;
        } else {
            min = p
        }

        if max - min <= 1 {
            // 从右至左，大到小，判断，使得，插入 n 后，
            // n (v[p]) < v[p + 1]
            if v[max] <= n {
                return max + 1
            } else if v[min] <= n {
                return min + 1
            } else {
                return min
            }
        }
    }
}


// 从 0 到 1024，随机生成一个长度小于 10 的 vector
fn crate_range_vector() -> Vec<i32> {
    let mut v: Vec<i32> = Vec::new();
    // get size first
    let mut rng = thread_rng();
    let n = rng.gen_range(0..10);
    if n == 0 {
        return v
    }

    let mut key = 0;
    key = rng.gen_range(key..1024);

    v.push(key);

    for _ in 1..n {
        key = rng.gen_range(key..1024);
        v.push(key);
    }

    v
}
