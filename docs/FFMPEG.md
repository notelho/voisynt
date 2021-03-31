# Ffmpeg
A list of the application ffmpeg commands.

## Make thick voice
```bash
ffmpeg -i input.mp3 -af "asetrate=44100*0.25,aresample=44100,atempo=2.17" output.mp3
```

## Make thin voice
```bash
ffmpeg -i input.mp3 -af "asetrate=44100*0.6,atempo=0.9" output.mp3
```

## Download
```bash
ffmpeg -i "http://streaming.input.foo/" -c copy output.mp3
```

## Make robotic
```bash
ffmpeg -i input.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" robot-thick.mp3
```

## Merge with channels and volume adjusts
```bash
ffmpeg -i input.mp3 -i input.mp3 -filter_complex "[0:a]volume=0.5[a1];[1:a]volume=1[a2];[a1][a2]amerge" -c:v copy -shortest output.mp3
```

## Merge files
```bash
ffmpeg -i input.mp3 -i input.mp3 -i input.mp3 -filter_complex "amix=inputs=3:duration=first:dropout_transition=3" output.mp3
ffmpeg -i input.mp3 -i input.mp3 -filter_complex "amix=inputs=2:duration=first:dropout_transition=2" output.mp3
```

## Change stereo to mono audio
```bash
ffmpeg -i input.mp3 -ac 1 output.mp3
```

## Concat files
```bash
ffmpeg -i "concat:input.mp3|input.mp3|input.mp3" -acodec copy output.mp3
```

## See audio info 
```bash
ffmpeg -i input.mp3
```