package curses

import (
    "errors"
    "fmt"
    "github.com/rthornton128/goncurses"
)

func InitCurses () (err error) {
    stdscr, err := goncurses.Init()

    if err != nil {
        etxt := fmt.Sprintf("go-curses.InitCurses/*/%s", err)
        err = errors.New(etxt)
        return
    }

    defer goncurses.End()

    goncurses.CBreak(true)
    goncurses.Echo(false)

    stdscr.Keypad(true)
    stdscr.Refresh()

    return
}

type Curses struct {
    // The curses window structure
    win goncurses.Window

    YMax int
    XMax int

    //  Size of the window
    height int
    width int

    //  Current position of the window on the screen
    posY int
    posX int

    //  Current cursor position within the window
    cursorY int
    cursorX int
}

func CursesWindow (height, width, starty, startx int) (cw Curses, err error) {
 //   var window goncurses.Window

    cw.height = height
    cw.width = width
    cw.posY = starty
    cw.posX = startx

    window, err := goncurses.NewWindow(height, width, starty, startx)
    cw.win = *window

    if err != nil {
      etxt := fmt.Sprintf("curses.CursesWindow/*/%s", err)
      err = errors.New(etxt)
      return
    }

    err = cw.DrawBorder(goncurses.ACS_VLINE, goncurses.ACS_HLINE)

    if err != nil {
      etxt := fmt.Sprintf("curses.CursesWindow/*/%s", err)
      err = errors.New(etxt)
    } else {
      cw.YMax, cw.XMax = cw.GetMaxYX()

      // Turn the cursor OFF
      SetCursor(0)
    }

    return
}

func (c *Curses) ClearWindow () (err error) {
    err = c.win.Clear()
    SetCursor(0)

    return
}

func (c *Curses) DrawBorder (vch, hch goncurses.Char) (err error) {
    err = c.win.Box(vch, hch)

    if err != nil {
      etxt := fmt.Sprintf("curses.DrawBorder/*/%s", err)
      err = errors.New(etxt)
    } else {
      c.win.Refresh()
    }

    return
}

func (c *Curses) GetMaxYX () (mY, mX int) {
  mY = c.height - 1
  mX = c.width - 1

  c.YMax = mY
  c.XMax = mX

  return
}

func (c *Curses) GetSize () (h, w int) {
    h = c.height
    w = c.width

    return
}

func (c *Curses) GetYX () (y, x int) {
    y, x = c.win.YX()

    return
}

func (c *Curses) GotoYX (y, x int) {
    c.win.Move(y, x)
    c.UpdateCursor(y, x)
}

func (c *Curses) UpdateCursor (y, x int) {
    c.cursorY = y
    c.cursorX = x
}

func (c *Curses) WriteStringAt (y, x int, s string, ref bool) {
    c.win.MovePrint(y, x, s)
    c.UpdateCursor(y, x)

    if ref {
        c.win.Refresh()
    }
}

func (c *Curses) WriteStringCentered (y int, s string, ref bool) {
    x := (c.width - len(s)) / 2

    c.WriteStringAt(y, x, s, ref)

    if ref {
        c.win.Refresh()
    }
}

func End () {
  goncurses.End()
}

func SetCursor (vis byte) (err error) {
    if vis < 0 || vis > 2 {
      etxt := fmt.Sprintf("curses.SetCursor/*/%s", err)
      err = errors.New(etxt)
    } else {
      goncurses.Cursor(vis)
    }

    return
}
