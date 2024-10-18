import { describe, expect, test } from "vitest";
import { getInputByDay } from "../utils";

import { first, second } from "./01";

const testInput = getInputByDay("01", "sample.txt");
const testInputTwo = getInputByDay("01", "sample-02.txt");
const sampleInput = getInputByDay("01");

describe("trebuchet", () => {
  test("first case", () => {
    const results = first(testInput);

    expect(results).toBe(142);

    console.log("01/01", first(sampleInput));
  });

  test.only("second case", () => {
    const results = second(testInputTwo);

    expect(results).toBe(281);

    console.log("01/02", second(sampleInput));
  });
});
