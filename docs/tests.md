-filter_complex "[0]volume=0.5,pan=2c[a];[1]volume=0.7,pan=2c[b];[a][b]amix=duration=shortest"

ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter_complex "amix=inputs=2:duration=first:dropout_transition=2,[0]volume=0.6;[1]volume=0.4" OUTPUT.mp3

ffmpeg -i thick-thin-robot.mp3 -i robot.mp3 -filter_complex "[1:a]volume=0.6,apad[A];[0:a][A]amerge[out]" -c:v copy -map 0:v -map [out] -y output-final.mp4

ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter:a "volume=0.5" -filter_complex "amix=inputs=2:duration=first:dropout_transition=2" OUTPUT.mp3

ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter_complex "[0:a]volume=1.6[a0]; [1:a]volume=0.4[a1]; [a0][a1]amix=inputs=2[a]" outpica.mp3

ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter_complex "[0:a]volume=0.5[a1];[1:a]volume=0.4[a2];[a1][a2]amerge" -c:v copy -shortest output.mp3

ffmpeg -i thick.mp3 -i robot.mp3 -filter_complex amix=inputs=2:duration=first:dropout_transition=2 TATA.mp3

ffmpeg -i robot.mp3 -i thick-thin.mp3 -filter_complex "[0:a]volume=3[a1];[1:a]volume=0.5[a2];[a1][a2]amerge" -c:v copy -shortest OUT.mp3

ffmpeg -i TATA.mp3 -i thin.mp3 -filter_complex "[0:a]volume=5[a1];[1:a]volume=0.1[a2];[a1][a2]amerge" -c:v copy -shortest BLOW.mp3

ffmpeg -i robot.mp3 -i thick-thin.mp3 -filter_complex "[0:a]volume=3[a1];[1:a]volume=0.5[a2];[a1][a2]amerge" -c:v copy -shortest OUT.mp3


ffmpeg -i robot.mp3 -i thick-thin-robot.mp3 -filter_complex "[0:a]volume=0.5[a1];[1:a]volume=1[a2];[a1][a2]amerge" -c:v copy -shortest output.mp3


ffmpeg -i thick.mp3 -i thin.mp3 -filter_complex "[0:a]volume=1[a1];[1:a]volume=0.5[a2];[a1][a2]amerge" -c:v copy -shortest thick-thin.mp3
ffmpeg -i thick.mp3 -i thin.mp3 -filter_complex "[0:a]volume=1[a1];[1:a]volume=0.5[a2];[a1][a2]amerge" -c:v copy -shortest thick-thin.mp3

ffmpeg -i robot.mp3 -i thick-thin.mp3 -filter_complex "[0:a]volume=2[a1];[1:a]volume=0.5[a2];[a1][a2]amerge" -c:v copy -shortest final.mp3