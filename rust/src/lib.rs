use suborbital::runnable::*;

struct Rust{}

impl Runnable for Rust {
    fn run(&self, input: Vec<u8>) -> Result<Vec<u8>, RunErr> {
        let in_string = String::from_utf8(input).unwrap();
        let n = in_string.parse::<i32>().unwrap();
        let f = factorial(n);
    
        Ok(String::from(format!("factorial of {} is {}", in_string, f)).as_bytes().to_vec())
    }
}

fn factorial(n: i32) -> i32 {
    if n == 1 { return 1 }

    return n * factorial(n-1)
}

// initialize the runner, do not edit below //
static RUNNABLE: &Rust = &Rust{};

#[no_mangle]
pub extern fn _start() {
    use_runnable(RUNNABLE);
}
