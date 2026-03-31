const increment = (x: number) => x + 1;
const concat = (a: string, b: string) => a + b;
const push = (arr: number[], value: number) => arr.push(value);
const pushReplacing = (arr: number[], value: number) => arr = [...arr, value];
const add = (obj: Record<string, string>, key: string, value: string) => obj[key] = value;
const addReplacing = (obj: Record<string, string>, key: string, value: string) => obj = { ...obj, [key]: value };

function performIncrement() {
    console.log("increment:")
    let a = 0
    increment(a)
    console.log(`a=${a}`)
    console.log("========================================")
}

function performConcat() {
    console.log("concat:")
    const abc = "abc"
    concat(abc, "def")
    console.log(`abc=${abc}`)
    console.log("========================================")
}

function performPush() {
    console.log("push:")
    const arr = [1, 2, 3]
    push(arr, 4)
    console.log(`arr=${arr}`)
    console.log("========================================")
}

function performPushReplacing() {
    console.log("pushReplacing:")
    const arr2 = [10, 20, 30]
    pushReplacing(arr2, 40)
    console.log(`arr2=${arr2}`)
    console.log("========================================")
}


function performAdd() {
    console.log("add:")
    const obj = { apple: "fruit", tomato: "fruit" }
    add(obj, "potato", "vegetable")
    console.log(`obj=${JSON.stringify(obj)}`)
    console.log("========================================")
}

function performAddReplacing() {
    console.log("addReplacing:")
    const obj2 = { car: "vehicle", bike: "vehicle" }
    addReplacing(obj2, "bus", "vehicle")
    console.log(`obj2=${JSON.stringify(obj2)}`)
    console.log("========================================")
}

performIncrement();
performConcat();
performPush();
performPushReplacing();
performAdd();
performAddReplacing();

// Pergunta: por que isso acontece em NodeJS?
// https://www.typescriptlang.org/play/?#code/MYewdgzgLgBAlmYAnApgWxWWBeGAKADwC4YwBXNAIxSQEoZsA+GAmAahgEYBuAKF4A2KWAEMGMAAy8EydJih4RtXqEgghAOgEgA5ngAGI7ABIA3iIC++5SvDQYq4CJz4RJaEgQ6ANDEruoTzAdeiYYMQ5KPltIUUpgcQAiEXjEmKcFFOBfRIATFAAzROVVCHUULV0DLJNzeKsbGPsABzIIAAtxRSQkEnIqGgBtAF1fADcRATIUPopqOgZmER6NVo68CamUWmjS0R7xQc5fACZfAGZh3jX27qRfABYSu3LKvUMe2uWkBv49mBuACUUM0BCJgF4ut9ZgMkCNxpNpjD5qElgdcIMNFjvgitsNdnZ9kgTodOBJTuSYOcJFcgSCwRDgnczjAHhJnmpNNp3t8Tl8eidfk1RLlcl0QJQAFYkYGgJC5AA8Hi8vmVwUYvgA1igAJ4BII+GCbJEwNUhRYwCWSwbanXDcTGlAE2KWqXiUzhZqgmYwRIFJBkOBQRK+KAgNDOEAkP0BoOJGAWXgiUV4K05ZogKCRkO+sYoHTCFJCYpNV7cgxW2oAKQAygB5AByGjNcAKOtTUtoQv+ydywNB4MhuA70pgspA8qVgRVpun6q1uv1M8dS+CqNdkvdMCxGjTMBtuuGJEdCed9itJNwHqcvVzKHacGAxd8lDg2ujeYfT5Q8cTvf7DJeCOLKJJQbQ5okn6PsWHJlFyVT6Be1b1k2LZtsBXbWPwAD02EwAACjQOhkFgbgAhOMAAI7TPAEBlOEqhQCgwAoDA6AwA2ID5LWAD8vBAA