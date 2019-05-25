# How to prepare sound with audacity

Download audacity [here](https://www.audacityteam.org/download/)

Open any sound you like, for [instance](https://freesound.org/people/mboscolo/sounds/212663/)
![ac-01](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_01.png?raw=true)

If your sample is a stereo - just split stereo to mono
![ac-02](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_02.png?raw=true)

And delete unused track
![ac-03](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_03.png?raw=true)

Select no more than one second of sample you like to install to Ninebot Zxx and ctrl-c it (copy to clipbuffer)
![ac-04](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_04.png?raw=true)

Press ctrl-n (new file) and ctrl-p (past from clipbuffer) sample we copying before
![ac-05](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_05.png?raw=true)

Add some effects if you wish so, at least amplification eventually needed (Effects->Amplification) or gain increase
![ac-06](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_06.png?raw=true)

After everything is ready - reduce sample rate to 16kHz
![ac-06](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_07.png?raw=true)

Then export audio
![ac-06](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_08.png?raw=true)

As raw audio signed 16 bit. Size of file should not exceed 32767 bytes, if not - you did something wrong
![ac-06](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_09.png?raw=true)

And process it with wnb-slicer
![ac-06](https://github.com/berghauz/wnb-slicer/blob/master/howto/images/as_10.png?raw=true)
