const name = process.argv[2];
console.log(process.argv);
const location = process.argv[3];
console.log(`Hi, I'm ${name}! I live in ${location}`);
console.log("process.env:", process.env);
console.log("process.env.NODE_ENV: ", process.env.NODE_ENV);