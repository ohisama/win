package main
import (
	"syscall"
	"ohisama/win"
)
var x int32
func WndProc(hWnd syscall.Handle, msg uint32, wParam, lParam uintptr) (uintptr) {
	switch msg {
	case win.WM_PAINT:
		x++
		var strMessage = "Hello"
		var ps win.PAINTSTRUCT
		hdc := win.BeginPaint(hWnd, &ps)
		for i := 0; i < 256; i++ {
			for j := 0; j < 256; j++ {
				iro := i * 256 * 256 + j
				win.SetPixelV(hdc, int32(i), int32(j), int32(iro))
			}
		}
		win.Ellipse(hdc, x,  200, 200 + x, 400)
		win.TextOut(hdc, 100, 100, strMessage, int32(len(strMessage)))
		win.EndPaint(hWnd, &ps)
		return 0
	case win.WM_COMMAND:
		return 0
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
		return 0
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}
func main() {
	win.WinMain(syscall.NewCallback(WndProc))
	return
}
