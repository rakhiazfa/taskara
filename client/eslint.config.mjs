import { FlatCompat } from "@eslint/eslintrc";
import importPlugin from "eslint-plugin-import";
import prettierPlugin from "eslint-plugin-prettier";
import { dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  ...compat.extends("next/core-web-vitals", "next/typescript"),
  {
    files: ["src/**/*.{js,jsx,ts,tsx}"],
    ignores: ["node_modules", ".next", "dist"],
    plugins: {
      prettier: prettierPlugin,
      import: importPlugin,
    },
    rules: {
      "prettier/prettier": [
        "error",
        {
          endOfLine: "auto",
        },
      ],
      "import/order": [
        "error",
        {
          groups: [
            "external",
            "builtin",
            "internal",
            "sibling",
            "parent",
            "index",
          ],
          alphabetize: {
            order: "asc",
            caseInsensitive: true,
          },
          "newlines-between": "always",
        },
      ],
    },
  },
];

export default eslintConfig;
