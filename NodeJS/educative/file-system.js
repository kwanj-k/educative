const fs = require("fs").promises;

const getStats = async path => {
    try {
        const stats = await fs.stat(path);
        console.log(stats);
        console.log(`isFile: ${stats.isFile()}`);
        console.log(`isDirectory: ${stats.isDirectory()}`);

    } catch (error) {
        console.error(error);
    }
};

// getStats('./test.txt');

// fs.writeFile("./test0.txt", "hello world");

const appendFile = async (path, data) => {
    try {
        await fs.appendFile(path, data);
    } catch (error) {
        console.error(error);
    }
};
appendFile("./test0.txt", "appending another hello world");

const readFile = async (path) => {
    try {
        const contents = await fs.readFile(path, "utf8");
        console.log(contents);
    } catch (error) {
        console.error(error);
    }
};

readFile("./test0.txt");