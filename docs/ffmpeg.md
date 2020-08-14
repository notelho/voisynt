# Commands

## Used commands

### Make thick voice
```bash
ffmpeg -i test.mp3 -af "asetrate=44100*0.25,aresample=44100,atempo=2.17" thick.mp3
```

### Make thin voice
```bash
ffmpeg -i test.mp3 -af "asetrate=44100*0.6,atempo=0.9" thin.mp3
```

### Download
```bash
ffmpeg -i "http://streaming.foo.com/" -c copy output.mp3
```

### Make robotic
```bash
ffmpeg -i thick.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" robot-thick.mp3
```

### Merge with channels and volume adjusts
```bash
ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter_complex "[0:a]volume=0.5[a1];[1:a]volume=1[a2];[a1][a2]amerge" -c:v copy -shortest output.mp3
```

### Merge files
```bash
ffmpeg -i thick.mp3 -i thin.mp3 -i robot.mp3 -filter_complex amix=inputs=3:duration=first:dropout_transition=3 thick-thin.mp3
ffmpeg -i thick.mp3 -i thin.mp3 -filter_complex amix=inputs=2:duration=first:dropout_transition=2 thick-thin.mp3
```

### Change stereo to mono audio
```bash
ffmpeg -i final.mp3 -ac 1 mono.mp3
```

## Unused commands

### Make loudnorm

ffmpeg -i thick.mp3 -af loudnorm=I=-16:TP=-1.5:LRA=14 -ar 48k loud.mp3

### Super equalizer

ffmpeg -i test.mp3 -af "superequalizer=1b=10:2b=10:3b=1:4b=5:5b=7:6b=5:7b=2:8b=3:9b=4:10b=5:11b=6:12b=7:13b=8:14b=8:15b=9:16b=9:17b=10:18b=10[a];[a]loudnorm=I=-16:TP=-1.5:LRA=14" -ar 48k super.mp3

### Concat files

ffmpeg -i "concat:thin2.mp3|thick.mp3|robotzada.mp3" -acodec copy concat_out.mp3

### See audio info 

ffmpeg -i test.mp3
