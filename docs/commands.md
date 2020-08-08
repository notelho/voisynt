# Commands

## Make thick voice

```bash
ffmpeg -i test.mp3 -af asetrate=44100*0.25,aresample=44100,atempo=2.6 thick.mp3
```

## Make thin voice

```bash
ffmpeg -i super.mp3 -af "asetrate=44100*1.2,atempo=1.5" thin.mp3
```

## Make loudnorm

ffmpeg -i thick.mp3 -af loudnorm=I=-16:TP=-1.5:LRA=14 -ar 48k loud.mp3

## Make robotic

ffmpeg -i super.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" robot3.mp3

## Super equalizer

ffmpeg -i test.mp3 -af "superequalizer=1b=10:2b=10:3b=1:4b=5:5b=7:6b=5:7b=2:8b=3:9b=4:10b=5:11b=6:12b=7:13b=8:14b=8:15b=9:16b=9:17b=10:18b=10[a];[a]loudnorm=I=-16:TP=-1.5:LRA=14" -ar 48k super.mp3

## Concat files

ffmpeg -i "concat:thin2.mp3|thick.mp3|robotzada.mp3" -acodec copy concat_out.mp3

## Merge files

ffmpeg -i thick.mp3 -i thin2.mp3 -i robotzada.mp3 -filter_complex amix=inputs=3:duration=first:dropout_transition=3 OUTPUT.mp3