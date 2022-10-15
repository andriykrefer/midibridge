package midi_parser

type BufferParser struct {
	buf []byte
}

func (thiss *BufferParser) Process(in []byte) (msgs [][]byte) {
	if len(in) == 0 {
		return
	}
	for k := range in {
		if isStatusByte(in[k]) {
			// Finish previous message
			if isMsgValid(thiss.buf) {
				msgs = append(msgs, thiss.buf)
			}
			thiss.buf = []byte{}
		}
		thiss.buf = append(thiss.buf, in[k])
		isLastByte := (k == len(in)-1)
		if isLastByte && isMidiMsgComplete(thiss.buf) {
			msgs = append(msgs, thiss.buf)
			thiss.buf = []byte{}
		}
	}
	return
}

func isMsgValid(msg []byte) bool {
	if len(msg) == 0 {
		return false
	}
	return isStatusByte(msg[0])
}

// Running status: http://midi.teragonaudio.com/tech/midispec.htm
// msgLen: Full midi message length
// dataLen: Data part length of a single message
func isRunningStatusLenOk(msgLen, dataLen int) bool {
	if msgLen < 1 {
		return false
	}

	dataLenNorm := dataLen
	if dataLen == 0 {
		dataLenNorm = 1
	}

	return (((msgLen - 1) % dataLenNorm) == 0)
}

func isMsgStatusAndLenOk(msg []byte, startRange, endRange byte, dataLen int) bool {
	isMsgLenOk := len(msg) == (dataLen + 1)
	isLenOk := isRunningStatusLenOk(len(msg), dataLen) || isMsgLenOk
	return isLenOk &&
		msg[0] >= startRange &&
		msg[0] <= endRange
}

func isSysexMsgComplete(msg []byte) bool {
	return msg[0] == 0xF0 &&
		msg[len(msg)-1] == 0xF7
}

func isMidiMsgComplete(msg []byte) bool {
	switch {
	case isMsgStatusAndLenOk(msg, 0xB0, 0xBF, 2): // Controller
	case isMsgStatusAndLenOk(msg, 0x80, 0x8F, 2): // Note off
	case isMsgStatusAndLenOk(msg, 0x90, 0x9F, 2): // Note on
	case isMsgStatusAndLenOk(msg, 0xA0, 0xAF, 2): // Aftertouch
	case isMsgStatusAndLenOk(msg, 0xB0, 0xBF, 2): // Controller
	case isMsgStatusAndLenOk(msg, 0xC0, 0xCF, 1): // Program Change
	case isMsgStatusAndLenOk(msg, 0xD0, 0xDF, 1): // Channel Pressure
	case isMsgStatusAndLenOk(msg, 0xE0, 0xEF, 2): // Pitch Wheel
	case isMsgStatusAndLenOk(msg, 0xF1, 0xF1, 1): // MTC Quarter Frame Message
	case isMsgStatusAndLenOk(msg, 0xF2, 0xF2, 2): // Song Position Pointer
	case isMsgStatusAndLenOk(msg, 0xF3, 0xF3, 1): // Song Select
	case isMsgStatusAndLenOk(msg, 0xF6, 0xF6, 0): // Tune Request
	case isMsgStatusAndLenOk(msg, 0xF8, 0xF8, 0): // MIDI Clock
	case isMsgStatusAndLenOk(msg, 0xF9, 0xF9, 0): // Tick
	case isMsgStatusAndLenOk(msg, 0xFA, 0xFA, 0): // Start
	case isMsgStatusAndLenOk(msg, 0xFC, 0xFC, 0): // Stop
	case isMsgStatusAndLenOk(msg, 0xFB, 0xFB, 0): // Continue
	case isMsgStatusAndLenOk(msg, 0xFE, 0xFE, 0): // Active Sense
	case isMsgStatusAndLenOk(msg, 0xFF, 0xFF, 0): // Reset
	case isSysexMsgComplete(msg): // System Exclusive
		return true
	}
	return false
}

func ProcessCompleteBlock(in []byte) (msgs [][]byte) {
	if len(in) == 0 {
		return
	}
	lastStatusByteIx := 0
	for k := range in {
		if isStatusByte(in[k]) {
			if k == 0 {
				continue
			}
			// Finish previous message
			m := in[lastStatusByteIx:k]
			msgs = append(msgs, m)
			lastStatusByteIx = k
		}

		isLastByte := (k == len(in)-1)
		if isLastByte {
			m := in[lastStatusByteIx : k+1]
			msgs = append(msgs, m)
		}
	}
	return
}

func isStatusByte(b byte) bool {
	return (b & 0x80) != 0
}
