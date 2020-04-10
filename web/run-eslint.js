#!/usr/bin/env node
/* eslint-disable */

/**
 * Runs eslint on the current directory, plus the "core" and "react" directories.
 *
 * This script is needed because eslint won't load the module configuration from
 * package.json in the nested directories, so the "react" directory linting would
 * fail because it would not find the react package, for instance.
 *
 * This script accepts multiple filenames separated by commas (,), lints and fixes
 * them.
 *
 * Usage: ./run-eslint.js path/to/the/file.js,another/file.js
 */

// Validate correct script usage
if (typeof process.argv[2] === 'undefined' || process.argv[2].length == 0) {
    console.log('Usage: ./run-eslint.js comma-sepparated,list-of-files');
}

const path = require('path');
const exec = require('child_process').execSync;

const subprojects = ['core', 'react'];
const linters = {
  main: {
    workingDir: __dirname,
    files: [],
  },
  core: {
    workingDir: path.resolve(__dirname, 'core'),
    files: [],
  },
  react: {
    workingDir: path.resolve(__dirname, 'react'),
    files: [],
  },
};

// Loop through every file passed in the args, and separate them in their linters
process.argv[2].split(',').forEach((filename) => {
  for (let i = 0; i < subprojects.length; i++) {
    const subproject = subprojects[i];

    // We allow passing relative paths, and here we convert them for standardization
    if (!path.isAbsolute(filename)) {
      // pre-commit will pass us the files relative to the project root
      // (aka with a `web` prefix). When it happens, this will add a '../'
      // allowing us to support paths relative to the current directory too.
      directoryPrefix = __dirname.slice(-3) == 'web' ? '../' : '';

      filename = path.join(__dirname, directoryPrefix, filename);
    }

    if (filename.includes(path.sep + subproject + path.sep)) {
      linters[subproject].files.push(filename);

      return;
    }
  }

  linters.main.files.push(filename);
});

let exitCode = 0;

for (const key in linters) {
  const linter = linters[key];

  // Skip linters with empty files
  if (linter.files.length == 0) {
    continue;
  }

  const command = path.resolve(__dirname, './node_modules/.bin/eslint');
  const args = [
    `--fix`,
    `--resolve-plugins-relative-to=${linter.workingDir}`,
    linter.files.join(),
  ].join(' ');

  console.log('+', `${command} ${args}`);

  try {
    exec(`${command} ${args}`, { stdio: 'inherit' }, (err, stdout, stderr) => {
      if (err) {
        throw err;
      }
      if (stderr) {
        console.log(`stderr: ${stderr}`);
      }
      console.log(`stdout: ${stdout}`);
    });
  } catch (err) {
    console.log(`error: ${err.message}`);

    // Exiting anything other than 0 will make pre-commit fail
    exitCode = 1;
  }
}

process.exit(exitCode);
