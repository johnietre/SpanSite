#[no_mangle]
pub extern "C" fn Primer(n: i32) -> i32 {
  let mut prime: i32;
  let mut counter: i32 = 4;
  let mut i: i32 = 9;
  let mut j: i32;
  loop {
    prime = 1;
    j = i / 2;
    if j % 2 == 0 {
      j += 1
    }
    while j > 1 {
      if i % j == 0 {
        prime = 0;
        break;
      }
      j -= 2
    }
    counter += prime;
    if counter == n {
      return i
    }
    i += 2;
  }
}