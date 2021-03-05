package containerstack

import (
    "fmt"
)

type Containerstack struct {
    capacity    int
    memorySize  int
    defaultTime int
    timeCircle  int
}

type Containerstacker interface {
    init()
    MemoryControl(ms, id int)
    SetMemorySize(ms int) int
    TimeControl() (bool, int)
    Killing(id int)
    Hibernate()
    Close()
    Freezen()
}

func (c *Containerstack) init() {
    c.memorySize = 6
    c.capacity = 6
    c.timeCircle = 5
    c.defaultTime = 10
}

func (c *Containerstack) MemoryControl(ms, id int) {
    if c.capacity <= c.memorySize {
        fmt.Println("capacity less and equal than memory size")
        c.SetMemorySize(ms)
    } else {
        ok, id := c.TimeControl()
        if ok && id > 0 {
            c.Killing(id)
        } 
    }
}

func (c *Containerstack) SetMemorySize(ms int) int {
    c.memorySize = ms
    return c.memorySize 
}

func (c *Containerstack) TimeControl() (bool, int) {
    if c.timeCircle > c.defaultTime {
        return true, 1
    } else if c.timeCircle >= 2 * c.defaultTime {
        return true, 2
    } else if c.timeCircle == c.defaultTime {
        return true, 3
    }
    return false, 0
}

func (c *Containerstack) Killing(id int) {
    if id == 1 {
        c.Hibernate()
    } else if id == 2 {
        c.Close()
    } else {
        c.Freezen()
    }
}

func (c *Containerstack) Hibernate() {

}

func (c *Containerstack) Close() {

}

func (c *Containerstack) Freezen() {

}
