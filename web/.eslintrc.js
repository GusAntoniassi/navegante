module.exports = {
  env: {
    es6: true
  },
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
  },
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: 'module',
  },
  extends: [
    'airbnb/base',
    'plugin:@typescript-eslint/recommended',
    'prettier/@typescript-eslint',
    "plugin:import/errors",
    "plugin:import/warnings",
    "plugin:import/typescript",
  ],
  plugins: ['react', 'import', 'jsx-a11y'],
  rules: {
    'import/prefer-default-export': 'off',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/explicit-member-accessibility': 'off',
    'semi': ['error', 'always'],
    'comma-dangle': ['error', 'always-multiline'],
    'eol-last': ['error', 'always'],
  },
  settings: {
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts'],
    },
  },
};
