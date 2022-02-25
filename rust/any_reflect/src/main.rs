use std::any::Any;
use std::fmt::Debug;

fn get_type<T:Any+Debug>(value: &T) {
    let value = value as &dyn Any;

    match value.downcast_ref::<String>() {
        Some(_) => { println!("Value's type is String") }
        None => (),
    }

    match value.downcast_ref::<Vec<String>>() {
        Some(_) => { println!("Value's type is Vec<String>") }
        None => (),
    }
}
fn main() {
    println!("Hello, world!");
    let s = "string".to_string();
    let v = vec!["this".to_string(), "is".to_string(), "a".to_string(), "string".to_string()];

    get_type(&s);
    get_type(&v);
}
