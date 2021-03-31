# Voisynt
Voisynth is a golang command line program that transforms text into a robot voice.

## Requirements
This app requires ffmpeg and was tested with the version `git-2020-07-30-134a48a`. You can get this version [here](https://www.gyan.dev/ffmpeg/builds/packages/), but any similar version should work fine. You can get this ffmpeg release in the official [site](https://www.ffmpeg.org/download.html#releases) or in it's official [github](https://github.com/FFmpeg/FFmpeg). After downloading the binary executable in your environment variables.

> If you are a windows user, you can get the `git-2020-07-30-134a48a` version in [voisynt binaries of v1.0.0 version](https://github.com/notelho/voisynt/releases/tag/v1.0.0)

After the installation check with CLI command:

```bash
ffmpeg -version # ffmpeg version git-2020-07-30-134a48a [...]
```

## Usage
The command line interface is quite simple. Just provide the message that you want to generate an audio and the path in which you want to save the file, that the file containing the selected phrase will be created. Additionally, you can pass a temporary file path where the files are saved temporarily in the process of generating the final file.

Check example bellow:

```bash
voisynt 
    --message "hello, nice to meet you!"        # A string message
    --output "./dist"                           # The output dir
    --tmp "./temp"                              # The temp dir (optional)
```