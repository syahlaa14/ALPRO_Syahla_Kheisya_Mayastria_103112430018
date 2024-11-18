/*package main 
import “fmt” 
func main() { 
    var nam float64 
    var nmk string 
    fmt.Print(“Nilai akhir mata kuliah: “) 
    fmt.Scan(&nam) 
    if nam > 80 { 
        nam = “A” 
    } 
    if nam > 72.5 { 
        nam = “AB” 
    } 
    if nam > 65 { 
        nam = “B” 
    } 
    if nam > 57.5 { 
        nam = “BC” 
    } 
    if nam > 50 { 
        nam = “C” 
    } 
    if nam > 40 { 
        nam = “D” 
    } else if nam <= 40 { 
        nam = “E” 
    } 
    fmt.Println(“Nilai mata kuliah: “, nmk) 
} */

// a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah 
// eksekusi program tersebut sesuai spesifikasi soal?

// program tidak akan berjalan karena ada kesalahan dalam kode
// variabel nam bertipe float64, tetapi diisi dengan nilai string seperti "A", yang menyebabkan error
// kesalahan ini membuat program tidak sesuai spesifikasi dan tidak dapat berjalan dengan benar

// b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

// tipe data yang tidak sesuai, variabel nam bertipe float64, tetapi diisi dengan nilai string seperti "A", yang menyebabkan error
// variabel nmk dideklarasikan tetapi tidak pernah diisi nilai
// semua kondisi menggunakan if terpisah, sehingga dievaluasi tanpa urutan jelas, menyebabkan hasil akhir tidak sesuai
// program meminta input nam dalam bentuk angka desimal float64
// logika if-else if digunakan untuk mengevaluasi nilai berdasarkan rentang tertentu
// variabel nmk menyimpan nilai mata kuliah, yang kemudian ditampilkan sebagai output

// c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. 
// Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

package main

import "fmt"

func main() {
    var nam float64
    var nmk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah:", nmk)
}

