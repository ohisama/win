package win
import (
	"syscall"
	"unsafe"
	"fmt"
)
const (
	WS_OVERLAPPED = 0x00000000
	WS_POPUP = 0x80000000
	WS_CHILD = 0x40000000
	WS_MINIMIZE = 0x20000000
	WS_VISIBLE = 0x10000000
	WS_DISABLED = 0x08000000
	WS_CLIPSIBLINGS = 0x04000000
	WS_CLIPCHILDREN = 0x02000000
	WS_MAXIMIZE = 0x01000000
	WS_CAPTION = 0x00C00000
	WS_BORDER = 0x00800000
	WS_DLGFRAME = 0x00400000
	WS_VSCROLL = 0x00200000
	WS_HSCROLL = 0x00100000
	WS_SYSMENU = 0x00080000
	WS_THICKFRAME = 0x00040000
	WS_GROUP = 0x00020000
	WS_TABSTOP = 0x00010000
	WS_MINIMIZEBOX = 0x00020000
	WS_MAXIMIZEBOX = 0x00010000
	WS_TILED = WS_OVERLAPPED
	WS_ICONIC = WS_MINIMIZE
	WS_SIZEBOX = WS_THICKFRAME
	WS_TILEDWINDOW = WS_OVERLAPPEDWINDOW
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_POPUPWINDOW = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW = WS_CHILD
	WM_CREATE = 0x0001
	WM_DESTROY = 0x0002
	WM_PAINT = 0x000F
	WM_CLOSE = 0x0010
	WM_COMMAND = 0x0111
	COLOR_WINDOW = 5
	COLOR_BTNFACE = 15
	CS_VREDRAW = 0x0001
	CS_HREDRAW = 0x0002
	CW_USEDEFAULT = -2147483648
	SW_SHOWDEFAULT = 10
	BS_DEFPUSHBUTTON = 1
	PS_solid = 0
	PS_DASH = 1
	PS_DOT = 2
	PS_DASHDOT = 3
	PS_DASHDOTDOT = 4
	PS_NULL = 5
	PS_INSIDEFRAME = 6
)
type WNDCLASSEX struct {
	cbSize uint32
	style uint32
	lpfnWndProc uintptr
	cbClsExtra int32
	cbWndExtra int32
	hInstance syscall.Handle
	hIcon syscall.Handle
	hCursor syscall.Handle
	hbrBackground syscall.Handle
	lpszMenuName * uint16
	lpszClassName * uint16
	hIconSm syscall.Handle
}
type POINT struct {
	x uintptr
	y uintptr
}
type MSG struct {
	hWnd syscall.Handle
	message uint32
	wParam uintptr
	lParam uintptr
	time uint32
	pt POINT
}
type RECT struct {
	Left int32
	Top int32
	Right int32
	Bottom int32
}
type PAINTSTRUCT struct {
	hdc syscall.Handle
	fErace uint32
	rcPaint RECT
	fRestore uint32
	fIncUpdate uint32
	rgbReserved byte
}
var (
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")
	user32, _ = syscall.LoadLibrary("user32.dll")
	gdi32, _ = syscall.LoadLibrary("gdi32.dll")
	pGetModuleHandleW, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	pLoadIconW, _ = syscall.GetProcAddress(user32, "LoadIconW")
	pLoadCursorW, _ = syscall.GetProcAddress(user32, "LoadCursorW")
	pRegisterClassExW, _ = syscall.GetProcAddress(user32, "RegisterClassExW")
	pCreateWindowExW, _ = syscall.GetProcAddress(user32, "CreateWindowExW")
	pDefWindowProcW, _ = syscall.GetProcAddress(user32, "DefWindowProcW")
	pDestroyWindow, _ = syscall.GetProcAddress(user32, "DestroyWindow")
	pPostQuitMessage, _ = syscall.GetProcAddress(user32, "PostQuitMessage")
	pShowWindow, _ = syscall.GetProcAddress(user32, "ShowWindow")
	pUpdateWindow, _ = syscall.GetProcAddress(user32, "UpdateWindow")
	pGetMessageW, _ = syscall.GetProcAddress(user32, "GetMessageW")
	pTranslateMessage, _ = syscall.GetProcAddress(user32, "TranslateMessage")
	pDispatchMessageW, _ = syscall.GetProcAddress(user32, "DispatchMessageW")
	pSendMessageW, _ = syscall.GetProcAddress(user32, "SendMessageW")
	pPostMessageW, _ = syscall.GetProcAddress(user32, "PostMessageW")
	pBeginPaint, _ = syscall.GetProcAddress(user32, "BeginPaint")
	pEndPaint, _ = syscall.GetProcAddress(user32, "EndPaint")
	pTextOutW, _ = syscall.GetProcAddress(gdi32, "TextOutW")
	pSetPixelV, _ = syscall.GetProcAddress(gdi32, "SetPixelV")
	pEllipse, _ = syscall.GetProcAddress(gdi32, "Ellipse")
)
func GetModuleHandle(lpModuleName * uint16) (syscall.Handle) {
	ret, _, _ := syscall.Syscall(uintptr(pGetModuleHandleW), 1, uintptr(unsafe.Pointer(lpModuleName)), 0, 0)
	return syscall.Handle(ret)
}
func LoadIcon(instance syscall.Handle, iconname * uint16) (syscall.Handle) {
	ret, _, _ := syscall.Syscall(uintptr(pLoadIconW), 2, uintptr(instance), uintptr(unsafe.Pointer(iconname)), 0)
	return syscall.Handle(ret)
}
func LoadCursor(instance syscall.Handle, cursorname * uint16) (syscall.Handle) {
	ret, _, _ := syscall.Syscall(uintptr(pLoadCursorW), 2, uintptr(instance), uintptr(unsafe.Pointer(cursorname)), 0)
	return syscall.Handle(ret)
}
func RegisterClassEx(lpwcx * WNDCLASSEX) (uint16) {
	ret, _, _ := syscall.Syscall(uintptr(pRegisterClassExW), 1, uintptr(unsafe.Pointer(lpwcx)), 0, 0)
	return uint16(ret)
}
func CreateWindowEx(dwExStyle uint32, lpClassName * uint16, lpWindowName * uint16, dwStyle uint32, x int32, y int32, nWidth int32, nHeight int32, hWndParent syscall.Handle, hMenu syscall.Handle, hInstance syscall.Handle, lpParam uintptr) (syscall.Handle) {
	ret, _, _ := syscall.Syscall12(uintptr(pCreateWindowExW), 12, uintptr(dwExStyle), uintptr(unsafe.Pointer(lpClassName)), uintptr(unsafe.Pointer(lpWindowName)), uintptr(dwStyle), uintptr(x), uintptr(y), uintptr(nWidth), uintptr(nHeight), uintptr(hWndParent), uintptr(hMenu), uintptr(hInstance), uintptr(lpParam))
	return syscall.Handle(ret)
}
func DefWindowProc(hWnd syscall.Handle, Msg uint32, wParam uintptr, lParam uintptr) (uintptr) {
	ret, _, _ := syscall.Syscall6(uintptr(pDefWindowProcW), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	return uintptr(ret)
}
func PostQuitMessage(nExitCode int32) {
	syscall.Syscall(uintptr(pPostQuitMessage), 1, uintptr(nExitCode), 0, 0)
	return
}
func ShowWindow(hWnd syscall.Handle, nCmdShow int32) (bool) {
	ret, _, _ := syscall.Syscall(uintptr(pShowWindow), 2, uintptr(hWnd), uintptr(nCmdShow), 0)
	return bool(ret != 0)
}
func DestroyWindow(hWnd syscall.Handle) {
	syscall.Syscall(uintptr(pDestroyWindow), 1, uintptr(hWnd), 0, 0)
	return
}
func UpdateWindow(hWnd syscall.Handle) {
	syscall.Syscall(uintptr(pUpdateWindow), 1, uintptr(hWnd), 0, 0)
	return
}
func GetMessage(lpMsg * MSG, hWnd syscall.Handle, wMsgFilterMin uint32, wMsgFilterMax uint32) (int32) {
	ret, _, _ := syscall.Syscall6(uintptr(pGetMessageW), 4, uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax), 0, 0)
	return int32(ret)
}
func TranslateMessage(lpMsg * MSG) (bool) {
	r, _, _ := syscall.Syscall(uintptr(pTranslateMessage), 1, uintptr(unsafe.Pointer(lpMsg)), 0, 0)
	return bool(r != 0)
}
func DispatchMessage(lpMsg * MSG) (int32) {
	ret, _, _ := syscall.Syscall(uintptr(pDispatchMessageW), 1, uintptr(unsafe.Pointer(lpMsg)), 0, 0)
	return int32(ret)
}
func SendMessage(hWnd syscall.Handle, Msg uint32, wParam uintptr, lParam uintptr) (uintptr) {
	ret, _, _ := syscall.Syscall6(uintptr(pSendMessageW), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	return uintptr(ret)
}
func PostMessage(hWnd syscall.Handle, Msg uint32, wParam uintptr, lParam uintptr) {
	syscall.Syscall6(uintptr(pPostMessageW), 4, uintptr(hWnd), uintptr(Msg), uintptr(wParam), uintptr(lParam), 0, 0)
	return
}
func BeginPaint(hDC syscall.Handle, lpPaint * PAINTSTRUCT) (syscall.Handle) {
	ret, _, _ := syscall.Syscall(uintptr(pBeginPaint), 2, uintptr(hDC), uintptr(unsafe.Pointer(lpPaint)), 0)
	return syscall.Handle(ret)
}
func EndPaint(hDC syscall.Handle, lpPaint * PAINTSTRUCT) (syscall.Handle) {
	ret, _, _ := syscall.Syscall(uintptr(pEndPaint), 2, uintptr(hDC), uintptr(unsafe.Pointer(lpPaint)), 0)
	return syscall.Handle(ret)
}
func TextOut(hDC syscall.Handle, x int32, y int32, text string, cbString int32) (bool) {
	ret, _, _ := syscall.Syscall6(uintptr(pTextOutW), 5, uintptr(hDC), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), uintptr(cbString), 0)
	return bool(ret != 0)
}
func SetPixelV(hDC syscall.Handle, x int32, y int32, cbString int32) (int32) {
	ret, _, _ := syscall.Syscall6(uintptr(pSetPixelV), 4, uintptr(hDC), uintptr(x), uintptr(y), uintptr(cbString), 0, 0)
	return int32(ret)
}
func Ellipse(hDC syscall.Handle, x1 int32, y1 int32, x2 int32, y2 int32) (int32) {
	ret, _, _ := syscall.Syscall6(uintptr(pEllipse), 5, uintptr(hDC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), 0)
	return int32(ret)
}
func MakeIntResource(id uint16) (* uint16) {
	return (* uint16) (unsafe.Pointer(uintptr(id)))
}
func WinMain(wproc uintptr) {
	var IDC_ARROW = MakeIntResource(32512)
	var IDI_APPLICATION = MakeIntResource(32512)
	hInstance := GetModuleHandle(nil)
	lpszClassName := syscall.StringToUTF16Ptr("ohisama/win")
	lpszWindowName := syscall.StringToUTF16Ptr("ohisama/win")
	var wcex WNDCLASSEX
	wcex.cbSize = uint32(unsafe.Sizeof(wcex))
	wcex.style = CS_HREDRAW | CS_VREDRAW
	wcex.lpfnWndProc = wproc
	wcex.cbClsExtra = 0
	wcex.cbWndExtra = 0
	wcex.hInstance = hInstance
	wcex.hIcon = LoadIcon(hInstance, IDI_APPLICATION)
	wcex.hCursor = LoadCursor(0, IDC_ARROW)
	wcex.hbrBackground = COLOR_WINDOW + 1
	wcex.lpszMenuName = nil
	wcex.lpszClassName = lpszClassName
	wcex.hIconSm = LoadIcon(hInstance, IDI_APPLICATION)
	RegisterClassEx(&wcex)
	hWnd := CreateWindowEx(0, lpszClassName, lpszWindowName, WS_OVERLAPPEDWINDOW, CW_USEDEFAULT, CW_USEDEFAULT, 600, 600, 0, 0, hInstance, 0)
	ShowWindow(hWnd, SW_SHOWDEFAULT)
	UpdateWindow(hWnd)
	var msg MSG
	for {
		if GetMessage(&msg, 0, 0, 0) == 0 {
			break
		}
		TranslateMessage(&msg)
		DispatchMessage(&msg)
	}
	return
}
