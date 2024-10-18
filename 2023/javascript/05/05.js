import { getAllMatches } from "../utils";

export function first(input) {
  const { seeds, maps } = parser(input);

  const locations = [];
  for (const seed of seeds) {
    let num = seed;
    for (const map of maps) {
      num = map[num] || num;
    }
    locations.push(num);
  }

  return Math.min(...locations);
}

export function parser(input) {
  const [seeds, ...maps] = input.map((line) =>
    getAllMatches(line, /(\d+)/g).map((s) => Number(s))
  );

  return { seeds, maps: maps.map(generateMapping) };
}

function generateMapping(mapInput) {
  const mappings = {};

  let group = [];
  mapInput.forEach((num, i) => {
    group.push(num);
    switch (i % 3) {
      case 2: {
        const [destination, source, range] = group;

        for (let j = 0; j < range; j++) {
          mappings[source] = {
            destination,
            range,
          };
        }

        group = [];
      }
    }
  });

  return mappings;
}
