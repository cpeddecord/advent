export function parser(input) {
  return input.map(lineParser);
}

function lineParser(inputLine) {
  const [card, numberLists] = inputLine.split(":");

  const [winners, nums] = numberLists.split("|").map(mapper);

  function mapper(str) {
    return str
      .trim()
      .split(" ")
      .map((s) => {
        if (!s) return;
        return Number(s.trim());
      })
      .filter(Boolean);
  }

  return {
    card: Number(/\d+/.exec(card)[0]),
    winners: new Set(winners),
    nums: nums,
  };
}

export function first(input) {
  const cards = parser(input);
  let sum = 0;

  cards.forEach((c) => {
    const matched = c.nums.filter((n) => c.winners.has(n));

    if (matched.length === 0) {
      sum += 0;
    } else if (matched.length === 1) {
      sum += 1;
    } else {
      sum += Math.pow(2, matched.length - 1);
    }
  });

  return sum;
}

export function second(input) {
  const cards = parser(input);
  const cardNums = new Map();

  cards.forEach((c) => {
    const matched = c.nums.filter((n) => c.winners.has(n));

    if (matched.length) {
      const value = cardNums.get(c.card) || 0;
      if (value) {
      }

      cardNums.set(c.card, 1);
    }

    if (matched.length > 1) {
    }
  });
}
