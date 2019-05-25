# wnb-slicer

wnb-slicer is a tool that addressed to help in custom sound preparation process for your Ninebot Zxx series EUC.
Before use it you should prepare sound samples which meet the following requirements:
  - the file size is not bigger than 32700 bytes, multiple sequenced sounds still not supported if it is not fit in standard EUC sound mapping scheme
  - sound should be encoded as 16000Hz sample rate and signed 16bit PCM
  - one second of 16Khz/16bit give us 32000 bytes of data, it is common Ninebot Zxx sounds duration (except warning sounds, it could be 2 seconds long)
  - Magic

# TODO

  - resampling in place from wav
  - maybe a simple GUI

# How to use

Show options:
```sh
wnb-slicer --help
```
Slice one-page-in-memory raw sample:
```sh
wnb-slicer.exe -file-name samples/doc_who_themev2_01.raw
```
Slice two-pages-in-memory raw sample (still in WIP state):
```sh
wnb-slicer.exe -max-chunk-size -file-name samples/doc_who_themev4_full.raw
```

# How to prepare sound with audacity

Check it [here](https://github.com/berghauz/wnb-slicer/howto/)
