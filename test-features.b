extrn print;
extrn read;

factorial(n) {
    if (n <= 1) return (1);
    return (n * factorial(n - 1));
}

sum_to(n) {
    auto total, i;
    total = 0;
    i = 1;
    while (i <= n) {
        total =+ i;
        i =+ 1;
    }
    return (total);
}

add(a, b) {
    return (a + b);
}

abs_val(x) {
    if (x < 0)
        return (0 - x);
    else
        return (x);
}

main() {
    auto x, result;

    print(100);
    result = factorial(5);
    print(result);

    print(200);
    result = sum_to(10);
    print(result);

    print(300);
    result = add(7, 8);
    print(result);

    print(400);
    result = abs_val(0 - 5);
    print(result);

    print(500);
    x = 0;
    while (1) {
        if (x >= 5) break;
        print(x);
        x =+ 1;
    }

    print(600);
    x = 10;
    if (x > 20)
        print(1);
    else if (x > 5)
        print(2);
    else
        print(3);
}
