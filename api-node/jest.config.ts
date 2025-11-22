import type { Config } from "jest";

const config: Config = {
  preset: "ts-jest",
  testEnvironment: "node",
  roots: ["<rootDir>/tests"],
  modulePaths: ["<rootDir>/src"],
  moduleFileExtensions: ["ts", "js", "json"],
  verbose: true,
};

export default config;
