package constants

/*
Keypad                   Keyboard
+-+-+-+-+                +-+-+-+-+
|1|2|3|C|                |1|2|3|4|
+-+-+-+-+                +-+-+-+-+
|4|5|6|D|                |Q|W|E|R|
+-+-+-+-+       =>       +-+-+-+-+
|7|8|9|E|                |A|S|D|F|
+-+-+-+-+                +-+-+-+-+
|A|0|B|F|                |Z|X|C|V|
+-+-+-+-+                +-+-+-+-+
keyboard inputs will use raylib constants
*/

const (
	PIXEL_SIZE int32 = 20
	WIDTH      int32 = 64
	HEIGHT     int32 = 32
	KB_NULL    int32 = 0
	KB_1       int32 = 49
	KB_2       int32 = 50
	KB_3       int32 = 51
	KB_4       int32 = 52
	KB_Q       int32 = 81
	KB_W       int32 = 87
	KB_E       int32 = 69
	KB_R       int32 = 82
	KB_A       int32 = 65
	KB_S       int32 = 83
	KB_D       int32 = 68
	KB_F       int32 = 70
	KB_Z       int32 = 90
	KB_X       int32 = 88
	KB_C       int32 = 67
	KB_V       int32 = 86
)

// arrays cant be listed in constant namespace :((((
var KNOWN_KB = [...]int32{KB_X, KB_1, KB_2, KB_3, KB_Q, KB_W, KB_E, KB_A, KB_S, KB_D, KB_Z, KB_C, KB_4, KB_R, KB_F, KB_V}

var FONTSET = [...]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}
