import { getAllMatches } from "../utils";

export function first(input) {
  return finder(input).reduce((m, c) => {
    if (!c.adjacentChars.length) return m;

    return m + c.value;
  }, 0);
}

export function finder(input) {
  return input
    .map((line, rowNum) => {
      const matches = getAllMatches(line, /(\d+)/g);

      const mapped = matches.map((str) => {
        const index = line.indexOf(str);
        const len = str.length;

        const rows = [rowNum - 1, rowNum, rowNum + 1];
        const columns = Array.from(
          { length: len + 2 },
          (_, idx) => idx + index - 1
        );

        const adjacentChars = [];
        for (const row of rows) {
          for (const column of columns) {
            const maybeChar = input?.[row]?.[column];

            if (maybeChar && maybeChar !== "." && isNaN(Number(maybeChar))) {
              adjacentChars.push(maybeChar);
            }
          }
        }

        return {
          value: Number(str),
          adjacentChars,
        };
      });

      return mapped.flat();
    })
    .flat();
}
