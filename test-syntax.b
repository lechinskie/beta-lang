/* Multi-line block comment
   testing the skip rule */

extrn global_var;

w[7];
main() {
  auto a, b, c;
  auto str1, str2;

  a = 100;
  str1 = "Hello Lexer";
  str2 = 'Alternative String';

  /* Testing Keywords and Control Flow */
  if (a == 100) {
    b = a + 5;
  } else {
    b = a - 5;
  }
a:
label:
  192;

  while (b > 0) {
    --b;
    if (b != 0)
      goto label;
  }

  switch (a) {
  case 100:
    c = a * 2;
  case 50:
    c = a / 2;
  }
  // Test block statement
  {
    auto inner;
    inner = 42;
    c = inner;
  }
  a = *(w + 2);
  b = w[a];
  a = 2;
  w[7] = a + b;
  /* Testing Operators */
  c = (a << 2) | (b >> 1);
  c = a % 3;
  c = a & b;
  c = !a;

  if (a <= 100)
    if (b >= 0) {
    /* This is the B way to do "a && b"
     */}

  /* Testing Ternary and Comparisons */
  c = (a < b) ? a : b;

  return (c);
}

asm_f() {
  // for gas-x86_64-linux
  __asm__("movq $60, %rax", // exit syscall number
          "movq $69, %rdi", // exit code
          "syscall");
}

naked_f __asm__("movq $69, %rax", "ret");

// variadic usage
extrn printf;
__variadic__(printf, 1); // number of fixed args

main2() { printf("Hello, world!\n"); }
