
fn foo1(mut p1:i32, p2:i32, p3:i32) {
    return 1;
}

fn foo2() -> i32{
    return;
}

fn foo3() -> i32{
}

fn program_3_1__1(mut param:i32) {
    {
        foo1(1,1,1);
        foo1(1,1,1,1);
        foo1(1);
        foo1(foo1(1,1,1), 1, 1);
        param = foo2();
        param = foo2(4,1,2,param);
        param = foo1(1,2,3);
    }
    
    {
        let mut t:i32;
        t = 1 * 1 + 10 / 3 - 5 * 9 + param * foo1(t, t, t);
        param = t;
        return param;
        {
            {
                return t;
                return;
            }
        }
    }

    0;
    (1);
    ((2));
    (((3)));
    ;;;;

    let mut a;
    let b;

    b = 9;
    a = a + b;
    b = b * a - 1 / 90;
    191919191919191919191;
    
    {
        a;
        b;
        a;
        {
            
        }
    }

    while 1{
        let mut d ;
        d = 10;
    }

    if 1 {
        let mut e:i32;
        e = 20 * 3;
    }else{
        let f = 10;
        let mut e; 
        e = b * 3;
    }

    param = param - 1;

    program_3_1__1(1,2,3);
    return;
}

fn program_3_1__2(mut a:i32, b:i32, c:i32,mut d:i32, e:i32) -> i32{
    a;
    (a);
    ((a));
    (((a)));
    while 1{
        let varInWhile;
        loop {
            let varInLoop;

        }
        if a{

        } else{

        }
    }
    if 2 {

    };

    return;
    {
        return;
    };
    return;
}
