import { describe, expect, test } from "vitest";
import { getInputByDay } from "../utils";

import { first, second, parser } from "./02";

const sampleInput = getInputByDay("02", "sample.txt")
  .split("\n")
  .filter(Boolean);

const sampleInputTwo = getInputByDay("02", "sample-02.txt")
  .split("\n")
  .filter(Boolean);

const fullInput = getInputByDay("02").split("\n").filter(Boolean);

describe("parser", () => {
  test("parse line", () => {
    expect(parser(sampleInput[0])).toStrictEqual({
      id: 1,
      sets: [{ blue: 3, red: 4 }, { red: 1, green: 2, blue: 6 }, { green: 2 }],
    });
  });
});

describe("cube conundrum", () => {
  test("first", () => {
    const limits = { red: 12, green: 13, blue: 14 };
    const results = first(limits, sampleInput);

    expect(results).toBe(8);

    console.log("02/01", first(limits, fullInput));
  });

  test("second", () => {
    const limits = { red: 12, green: 13, blue: 14 };
    const results = second(sampleInputTwo);

    expect(results).toBe(2286);

    console.log("02/02", second(fullInput));
  });
});
