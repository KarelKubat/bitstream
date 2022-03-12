# bitstream

`bitstream` is a Go package to read or write bits. It has two sub-packages: `bitstream/bitrd` for reading, and `bitstream/bitwr` for writing.

- A `bitstream/bitrd` bit reader is instantiated with an underlying `io.Reader` that is called to fetch bytes. The bit reader then yields the bits using its method `Read()`.

- Similarly, a `bitstream/bitwr` bit writer is instantiated with an underlying `io.Writer`. The writer calletes bits that are sent to the bit writer using its method `Write()` into bytes, which then are sent to the underlying `io.Writer`.

Basically, the bitstream reader takes a standard `io.Reader` but returns 8 bits for every byte that this reader would yield. The bitstream takes a standard `io.Writer`, accepts bits, and for every 8 bits it has received, it sends a byte to its `io.Writer`.

The bit reader and writer don't know about any other datatype than a byte. There is no concept of runes, characters, ints etc.. Just bits, and groups of eight of them are bytes.

## Synopsis

```go
import "github.com/KarelKubat/bitstream"
import "github.com/KarelKubat/bitstream/bitrd"
import "github.com/KarelKubat/bitstream/bitwr"

func main() {
    // Any io.Reader from which bytes can be read
    stringReader := strings.NewReader("Hello World")

    // Bit reader
    rd := bitrd.New(stringreader)

    // Fetch the bits
    bits := []bitstream.Bit
    for {
        bit, err := rd.Read()
        switch {
        case err == nil:
            bits = append(bits, bit)
        case err == io.EOF:
            break
        case err != nil:
            log.Fatalln(err)
        }
    }

    // bits will now be (shown in groups of 8):
    // 01001000 01100101 01101100 01101100 01101111 00100000 
    // 01010111 01101111 01110010 01101100 01100100

    // Any io.Writer to which bytes can be sent
    stringWriter := new(strings.Builder)

    // Bit writer
    wr := bitwr.New(stringWriter)

    // Send the bits
    for _, bit := range bits {
        if err := wr.Write(bit); err != nil {
            log.Fatalln(err)
        }
    }

    // Here's the end result
    fmt.Println(stringWriter.String())
}
```

## `bitstream` API

Package `bitstream` only defines what a `Bit` is and has a complimentary `String()` function to display the bit as `0` or `1`.

- The type is `bitstream.Bit`.
- Value `bitstream.Zero` represents a 0-bit.
- Value `bitstream.One` represents a 1-bit.

## `bitstream/bitrd` API

- `r := New(rdr io.Reader`) returns an initialized bit reader. There is no automatic closing of the underlying `io.Reader`, just `close(r.Reader)` once appropriate.

- `r.Read()` returns a `bitstream.Bit` and an error. The error `io.EOF` indicates that there is no more to read, as with any standard reader. In this situation `r.Reader` may be closed.

## `bitstream/bitwr`

- `w := New(wtr io.Writer)` returns an initialized bit writer. There is no automatic closing of the underlying `io.Writer`, just `close w.Writer` once appropriate.
- `w.Write(bit bitstream.Bit)` writes one bit, which is either `bitstream.Zero` or `bitstream.One`. This method returns an error when the underlying `io.Writer` fails.
- `w.Flush()` flushes any incomplete byte to the underlying `io.Writer`. Calling `w.Flush()` is only required when the number of bits written by `w.Write()` isn't a multiple of eight.