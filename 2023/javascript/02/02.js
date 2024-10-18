export function parser(line) {
  const diceRegex = /(\d+) (\w+)/;

  const [game, setsStr] = line.split(":");
  const sets = setsStr.split(";");

  const parsedSets = sets.map((curr) => {
    return curr.split(",").reduce((m, s) => {
      const str = s.trim();
      const [, num, color] = diceRegex.exec(str);
      return { ...m, [color]: Number(num) };
    }, {});
  }, {});

  return {
    id: Number(/\d+/.exec(game)),
    sets: parsedSets,
  };
}

export function first(limits, inputLines) {
  const input = inputLines.map(parser);
  const validIds = [];

  input.forEach(({ id, sets }) => {
    const isPossible = sets.every((set) => {
      for (const color in set) {
        if (!limits[color]) {
          return false;
        }

        if (limits[color] < set[color]) {
          return false;
        }
      }

      return true;
    });

    if (isPossible) {
      validIds.push(id);
    }
  });

  return validIds.reduce((a, b) => a + b, 0);
}

export function second(inputLines) {
  const input = inputLines.map(parser);

  let sum = 0;

  input.forEach(({ sets }) => {
    const colorMaxes = {};
    sets.forEach((set) => {
      for (const color in set) {
        if (!colorMaxes[color] || colorMaxes[color] < set[color]) {
          colorMaxes[color] = set[color];
        }
      }
    });

    sum += Object.values(colorMaxes).reduce((a, b) => a * b);
  });

  return sum;
}
