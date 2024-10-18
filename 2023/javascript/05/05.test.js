import { describe, expect, test } from "vitest";
import { getInputByDay } from "../utils";

import { first, second, parser } from "./05";

const full = getInputByDay("05").trim().split(/^\s*$/m);
const testInput = getInputByDay("05", "sample.txt").trim().split(/^\s*$/m);
const testInput2 = getInputByDay("05", "sample-02.txt").trim().split("\\n\\s");

describe("Fertilizer", () => {
  test.only("first", () => {
    const results = first(testInput);

    expect(results).toBe(35);

    // console.log("05/01", first(full));
  });
});
