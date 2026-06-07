extrn print;
extrn read;

v[2];
main() {
  read(v[0]);
  read(v[1]);

  v[0] = v[0] + v[1];
  v[1] = v[0] - v[1];
  v[0] = v[0] - v[1];

  print(v[0]);
  print(v[1]);
  print(v);

  auto a;
  a = *(v);
  print(a);
  a = *(v + 4); // 4 é o tamanho do word
  print(a);
}
