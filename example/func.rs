fn path() -> i32{
    let mut a = 1;
    if a {
        return a;
    }
    return a - 1;
}

fn ifelse() {
    let mut a;
    a = 1;
    if a == 0 {
        a = a + 1;
        if a {
            a = a - 1;
        }
    }else {
        a = a * a;
    }

    a = a * 9;
    (a);
}

fn foo() -> i32 {
    let mut a;
    a = 1;
    a = 2;
    let mut b;
    b = 3;
    b = a + 1;
    return 1;
}

fn relation() -> i32 {
    let mut a;
    a = 1;
    a = 1 < 2;
    let mut b;
    b = 1 != a;
    b = 2 >= a;
    b = 2 <= a;
    b = 2 < a;
    b = 2 > a;
    b = 2 == a;
    return (a + b) / a;
}

fn goo() -> i32 {
    let a;
    a = 1;

    {
        a = a + a;
        a = a * a - 1;
        {
            return a;
        }
    }

    return foo();
}

fn main() -> i32 {
    foo();
    return 0;
}
