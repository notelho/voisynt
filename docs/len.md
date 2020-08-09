
## test
Duration: 00:00:02.06, start: 0.000000, bitrate: 32 kb/s

## thick
Duration: 00:00:02.06, start: 0.000000, bitrate: 32 kb/s

ffmpeg -i test.mp3 -af asetrate=44100*0.25,aresample=44100,atempo=2.17 thick.mp3

## thin
Duration: 00:00:02.06, start: 0.000000, bitrate: 32 kb/s

ffmpeg -i test.mp3 -af "asetrate=44100*1.07,atempo=0.5" thin.mp3

# robot
Duration: 00:00:02.06, start: 0.000000, bitrate: 32 kb/s

ffmpeg -i test.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" robot.mp3

## output

## thick-thin
ffmpeg -i thick.mp3 -i thin.mp3 -filter_complex amix=inputs=2:duration=first:dropout_transition=2 thick-thin.mp3
ffmpeg -i thick-thin.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" thick-thin-robot.mp3

ffmpeg -i test.mp3 -i thick.mp3 -filter_complex amix=inputs=2:duration=first:dropout_transition=2 OUTPUT.mp3
ffmpeg -i thick-thin-robot.mp3 -i robot.mp3 -filter_complex amix=inputs=2:duration=first:dropout_transition=2 OUTPUT.mp3

