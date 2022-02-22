package utils

type Color int

// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	Black   Color = 0
	Red     Color = 1
	Green   Color = 2
	Yellow  Color = 3
	Blue    Color = 4
	Magenta Color = 5
	Cyan    Color = 6
	White   Color = 7
)

type Mode int

const (
	Normal  Mode = 0
	Light   Mode = 1
	Under   Mode = 4
	Flash   Mode = 5
	Reverse Mode = 7
	Hide    Mode = 8
)

const (
	SpNormal Mode = 0
	SpLight  Mode = 1
	SpUnder  Mode = 4
	SpFlash  Mode = 5
	SpFill   Mode = 7
	SpHide   Mode = 8
)
