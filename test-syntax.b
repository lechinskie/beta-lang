/* Multi-line block comment
   testing the skip rule */

extrn print;
extrn read;

w[7];

main() {
  auto a, b, c;
	read(c);
  a = c;


	w[2] = 3;
  b = w[a];
	print(b);
	w[4] = b + 5;
	print(w[4]);
}

