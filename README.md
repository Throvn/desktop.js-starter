# Desktop.js CLI

> [!NOTE]
> Only works on macOS using ARM processors right now.

## Installation

Since I've only tested on mac, I'm assuming you are on a mac with an M-series chip.

1. Go to the github releases page.
2. Download `djs` (it's the compiled binary)
3. Move the binary to ` /usr/local/bin/`
4. In the terminal run `djs help` to test if the executable is accessible.

### Build

1. Clone this repo
2. run `./run.sh` in the project root (you will be prompted for your password)
3. Enter your password. This is only needed to copy the go binary to `/usr/local/bin` to make it executable systemwide. If you don't like to give sudo permissions, just run `go build` and you are done. Note though, that some of the commands make use of the current directory, which is a bit annoying if you don't have the executable in your `/usr/local/bin` directory.

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

Needs to be started in the project root.
Currently works only on macOS.
Bundles the code together with the engine into an app bundle.
The app bundle will be created inside of project root.
