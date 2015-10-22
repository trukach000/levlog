package levlog

import (
    "os"
    "sync"
    "time"
)

type RotateWriter struct {
    lock     sync.Mutex
    filename string // should be set to the actual filename
    fp       *os.File
}

// Make a new RotateWriter. Return nil if error occurs during setup.
func NewRotateWrite(filename string) (*RotateWriter, error) {
    w := &RotateWriter{filename: filename}
    err := w.Rotate()
    if err != nil {
        return nil,err
    }
    return w, nil
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
    w.lock.Lock()
    defer w.lock.Unlock()
    return w.fp.Write(output)
}

// Perform the actual act of rotating and reopening file.
func (w *RotateWriter) Rotate() (err error) {
    w.lock.Lock()
    defer w.lock.Unlock()

    // Close existing file if open
    if w.fp != nil {
        err = w.fp.Close()
        w.fp = nil
        if err != nil {
            return
        }
    }
    // Rename dest file if it already exists
    _, err = os.Stat(w.filename)
    if err == nil {
        err = os.Rename(w.filename, w.filename+"."+time.Now().Format("2006_01_02_15_04_05"))
        if err != nil {
            return err
        }
    }

    // Create a file.
    w.fp, err = os.Create(w.filename)
    return err
}