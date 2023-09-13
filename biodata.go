package main

import (
	"fmt"
	"os"
)

type Teman struct {
    Nama      string
    Alamat    string
    Pekerjaan string
    Alasan    string
}

var temanKelas = []Teman{
    {"Thomas", "Alamat Thomas", "Pekerjaan Thomas", "Alasan Thomas"},
    {"Ihsan", "Alamat Ihsan", "Pekerjaan Ihsan", "Alasan Ihsan"},
    {"Advie", "Alamat Advie", "Pekerjaan Advie", "Alasan Advie"},
    {"Faisal", "Alamat Faisal", "Pekerjaan Faisal", "Alasan Faisal"},
    {"Rifal", "Alamat Rifal", "Pekerjaan Rifal", "Alasan Rifal"},
}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        fmt.Println("Usage: go run biodata.go [Nama Teman]")
        os.Exit(1)
    }

    namaTeman := args[0]

    for i, teman := range temanKelas {
        if teman.Nama == namaTeman {
            fmt.Printf("Data teman dengan nama %s (Absen No %d):\n", teman.Nama, i+1)
            fmt.Printf("Nama: %s\n", teman.Nama)
            fmt.Printf("Alamat: %s\n", teman.Alamat)
            fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
            fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
            return
        }
    }

    fmt.Printf("Tidak ada teman dengan nama %s ditemukan\n", namaTeman)
}
