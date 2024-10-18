export function first(input) {
  const lines = input.split("\n").filter(Boolean);

  const nums = lines.map((line) =>
    line
      .split(/[a-zA-Z]/)
      .filter(Boolean)
      .join("")
      .split("")
  );

  return nums.reduce((prev, curr) => {
    const str = `${curr[0]}${curr[curr.length - 1]}`;
    return prev + Number(str);
  }, 0);
}

const numTuple = [
  ["0", "zero"],
  ["1", "one"],
  ["2", "two"],
  ["3", "three"],
  ["4", "four"],
  ["5", "five"],
  ["6", "six"],
  ["7", "seven"],
  ["8", "eight"],
  ["9", "nine"],
];

const numsMap = numTuple.reduce((prev, [a, b], i) => {
  return {
    ...prev,
    [a]: i,
    [b]: i,
  };
}, {});

const chars = numTuple
  .map(([a, b]) => {
    return `${a}|${b}`;
  }, "")
  .join("|");

const regex = new RegExp(`(?=(${chars}))`, "g");

export function second(input) {
  const lines = input.split("\n").filter(Boolean);

  const numPairs = lines.map((s) => {
    const nums = [];

    let m = regex.exec(s);
    while (m !== null) {
      if (m.index === regex.lastIndex) {
        regex.lastIndex++;
      }

      m.forEach((match) => {
        if (match) nums.push(numsMap[match]);
      });

      m = regex.exec(s);
    }

    return nums;
  });

  return numPairs.reduce((prev, curr) => {
    const str = `${curr[0]}${curr[curr.length - 1]}`;
    return prev + Number(str);
  }, 0);
}
