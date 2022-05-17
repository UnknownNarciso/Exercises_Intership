let n = 7;
// Upside pyramid
for (let i = 1; i <= n; i+=2) {
  // printing spaces
  for (let j = 1; j <= n - i/2; j++) {
    process.stdout.write(' ')
  }
  // printing star
  for (let k = 0; k <  i ; k++) {
    process.stdout.write('*')
  }
  console.log();
}

// downside pyramid
for (let i = n - 2; i > 0 ; i-=2) {
  // printing spaces
  for (let j = 1; j < n - i/2; j++) {
     process.stdout.write(' ');
  }
  // printing star
  for (let  j = 0; j< i; j++) {
    process.stdout.write('*');
  }
  console.log();
}