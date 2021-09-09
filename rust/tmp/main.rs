use std::collections::HashSet;

fn run() {
    let total = 60;
    let couples = 4;
    let couple = [7, 8];
    let mut exists: HashSet<(u8, u8)> = HashSet::new();
    let mut demo: Vec<u8> = Vec::new();
    for i in 0..60 {
        demo.push(i);
    }

    for _ in 0..couples {
        for members in couple {
            // 新建一个分组
            let mut group = Vec::new();
            // 插入第一个参数时，一定成功 index 0
            let k = demo[0];
            group.push(k);
            demo.remove(0);

            let mut can_insert = true;
            // 开始插入其它值
            for i in 0..demo.len() {
                can_insert = true;
                let k = demo[i];
                // 与组中现有的所有元素组合，查看是否已经有同组的情况发生
                for member in &group {
                    if !exists.insert((*member, k)) {
                        can_insert = false;
                        break;
                    }
                }

                // 可以插入此值，
                // 并从原始数组中移除
                if can_insert {
                    group.push(k);
                    demo.remove(i);
                }

                if group.len() == members {
                    break;
                }
            }
        }
    }
}

fn main() {
    println!("hello");
}
