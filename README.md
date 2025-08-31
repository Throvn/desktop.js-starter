# Desktop.js CLI

This program sets up your development environment for developing Desktop.js apps.

- It takes your JS/TS sources, bundles them and runs them.
- It can also produce production executables.

If you don't know what's [desktop.js](https://github.com/Throvn/desktop.js), head over to that repo first.

## Usage

First you want to initialize a new project.
The following command will create the folder structure for you.

```shell
djs init <path>
# e.g. djs init ./my-project
```

**Parameters**

- path: Where you want to create the project. Needs to be a not yet existing or empty folder path.

The project structure will look as follows:

```
.internals/
        javascript/
        types.d.ts
        GUI.d.mts
assets/
        fonts/
source/
        index.tsx
.gitignore
tsconfig.json
djs.ini
```

Note the `.internals/` directory is for temporary files.
All contents inside are generated. **You should not use it and you should not care**!

But in case you do care:

- It includes `types.d.ts` to set up autocomplete in your editor (assuming it's VSCode).
- The output of the bundler are also saved in there under `javascript/`.

The `assets/` directory is bundled as is into your package.
Apart from the fonts directory, which **absolutely needs to be there** and is not allowed to have anything other than `.tff` fonts inside, you can freely choose the directory structure.

The `index.jsx` is the main entry point and cannot be changed.

Inside of the djs.ini you will find the following:

```ini
[project]
name = Example Project
starter-version = 0.1.0

[window]
width = 600
height = 300
```

This file exists for global configurations.
If you adjust the file, you need to restart `djs` (in case you were in `watch` mode).

---

```shell
djs watch <path?>
```

Needs to be started in the project root.
Observes the `source/` directory for changes and rebundles the output.
The rebundled changes are then loaded and displayed.

---

```shell
djs bundle
```

Needs to be started in teh project root.
Currently works only on macOS.
Bundles the code together with the engine into an app bundle.
The app bundle will be created inside of project root.
