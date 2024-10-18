import { describe, expect, test } from "vitest";
import { getInputByDay } from "../utils";

import { first, second, finder } from "./03";

const full = getInputByDay("03").trim().split("\n");
const testInput = getInputByDay("03", "sample.txt").trim().split("\n");

describe("Gear Ratios", () => {
  test("finds adjacent numbers", () => {
    const results = finder(testInput);

    expect(results[0]).toStrictEqual({
      value: 467,
      adjacentChars: ["*"],
    });
  });

  test("finds numbers", () => {
    const input = ["....10.02.39..."];
    expect(finder(input)).toStrictEqual([
      { value: 10, adjacentChars: [] },
      { value: 2, adjacentChars: [] },
      { value: 39, adjacentChars: [] },
    ]);
  });

  test("finds corners", () => {
    const input = ["*..", ".1.", "..."];

    const results = finder(input);

    expect(results[0]).toStrictEqual({
      value: 1,
      adjacentChars: ["*"],
    });
  });

  test("part nums sum", () => {
    const results = first(testInput);

    expect(results).toBe(4361);

    console.log("03/01", first(full));
  });
});
