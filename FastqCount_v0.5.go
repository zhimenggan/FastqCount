package main

import ( "bufio"; "fmt"; "os"; "strings"; "strconv")
// import "compress/gzip"
import gzip "github.com/klauspost/pgzip"


func Exit () {
    fmt.Println ("Usage: FastqCount  <input.fastq>  [phred]")
    fmt.Println ("  Phred default: 33;")
    fmt.Println ("  Output (tsv): Total Reads  Total Bases  N Bases  Q20  Q30  GC;") 
    fmt.Println ("  Note: \"pigz -dc *.fastq.gz | FastqCount -\" is recommended for gzipped file(s).")
    fmt.Println ()
    fmt.Println ("author: d2jvkpn")
    fmt.Println ("version: 0.5")
    fmt.Println ("release: 2018-01-09")
    fmt.Println ("project: https://github.com/d2jvkpn/FastqCount")
    fmt.Println ("lisence: GPLv3 (https://www.gnu.org/licenses/gpl-3.0.en.html)")
    os.Exit (0)
}


func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        Exit ()
    }

    var scanner *bufio.Scanner
    var fname string = os.Args[1]
    var phred int

    if len(os.Args) == 3 {
        phred, _ = strconv.Atoi (os.Args[2])
    } else { phred = 33 }

    if fname == "-" {
        scanner = bufio.NewScanner (os.Stdin)
    } else {
        file, err := os.Open (fname)

        if strings.HasSuffix (fname, ".gz") {
            gz, _ := gzip.NewReader (file)
            scanner = bufio.NewScanner (gz)
        } else { scanner = bufio.NewScanner (file) }

        if err != nil { os.Exit (1) } else { defer file.Close () }
    }

    var i, bases, q20, q30, gc, Nc int = 0, 0, 0, 0, 0, 0

    for scanner.Scan() {
        i += 1; text := scanner.Text ()

        if i%4 == 2 {
            bases += len(text)
            Nc += strings.Count (text, "N")
            gc += strings.Count (text, "G")  + strings.Count (text, "C")
        } else if i%4 == 0 {
            for _, c := range text {
                if int(c) - phred >= 20 { q20 += 1 } else { continue }
                if int(c) - phred >= 30 { q30 += 1 }
            }
        }
    }

    fmt.Println ("Total Reads\tTotal Bases\tN Bases\tQ20\tQ30\tGC")

    fmt.Printf ("%d (%.2f M)\t%d (%.2f G)\t%.2f%%\t%.2f%%\t%.2f%%\t%.2f%%\n", 
        i/4, float32 (i) / 4E+6,
        bases, float32 (bases) / 1E+9,
        float32 (Nc*100 / bases),
        float32 (q20*100 / bases),
        float32 (q30*100 / bases),
        float32 (gc*100 / bases))
}
