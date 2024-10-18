import { describe, expect, test } from "vitest";
import { getInputByDay } from "../utils";

import { first, second, parser } from "./04";

const full = getInputByDay("04").trim().split("\n");
const testInput = getInputByDay("04", "sample.txt").trim().split("\n");
const testInput2 = getInputByDay("04", "sample-02.txt").trim().split("\n");

describe("Scratchcards", () => {
  test("parser", () => {
    const results = parser(testInput);

    expect(results[0]).toStrictEqual({
      card: 1,
      winners: new Set([41, 48, 83, 86, 17]),
      nums: [83, 86, 6, 31, 17, 9, 48, 53],
    });
  });

  test("scoring", () => {
    const results = first(testInput);

    expect(results).toBe(13);

    console.log("04/01", first(full));
  });

  test("card sums", () => {
    const results = second(testInput);

    expect(results).toBe(30);
  });
});
