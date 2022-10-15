# midibridge
Simple UART to MIDI Converter, to be used with Arduino or any other device that sends MIDI messages over UART serial.

It is a very simple application that forwards MIDI messages from UART to a MIDI device.

This implementation is more a proof of concept and is in Alpha stage.

I actually only testes with MIDI CC messages, but it should work with every message listed on [this reference](http://midi.teragonaudio.com/tech/midispec.htm).

## More info

Reference Arduino implementation: https://andriykrefer.com/simple-arduino-midi-controller-1-2/

More info here: https://andriykrefer.com/simple-arduino-midi-controller-2-2/

# Intructions
This application forwards MIDI messages from UART to a MIDI device. However, it does not create the virtual MIDI port itself. For that, you can use [loopMIDI](https://www.tobias-erichsen.de/software/loopmidi.html).

1. Create a virtual port on loopMIDI

![alt text](https://github.com/andriykrefer/midibridge/raw/master/img/loopMIDI.jpg)

2. Run midibridge. You can download a pre-compiled binary from the [releases](https://github.com/andriykrefer/midibridge/releases) page.

3. Select the UART port

![alt text](https://github.com/andriykrefer/midibridge/raw/master/img/com.jpg)

4. Select the created MIDI port

![alt text](https://github.com/andriykrefer/midibridge/raw/master/img/midi.jpg)

5. Open the create MIDI port on your DAW

6. Done
