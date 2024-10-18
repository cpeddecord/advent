import { readFileSync } from "node:fs";
import { resolve } from "node:path";
import { cwd } from "node:process";

function fileToString(path) {
  return readFileSync(path).toString();
}

export function getInputByDay(day, filename) {
  const filenameString = Boolean(filename) ? filename : "input.txt";
  const fullpath = [cwd(), day, filenameString].join("/");
  return readFileSync(fullpath).toString();
}

export function getAllMatches(string, regex) {
  const matches = [];
  let m = regex.exec(string);
  while (m !== null) {
    if (m.index === regex.lastIndex) {
      regex.lastIndex++;
    }

    matches.push(m[0]);

    m = regex.exec(string);
  }

  return matches;
}
