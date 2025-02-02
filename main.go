package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Struktur data untuk stasiun kereta
type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

// Data statis untuk daftar stasiun (ditambah lebih banyak)
var stations = []Station{
	{ID: "GMR", Name: "Gambir", City: "Jakarta"},
	{ID: "PSE", Name: "Pasar Senen", City: "Jakarta"},
	{ID: "BD", Name: "Bandung", City: "Bandung"},
	{ID: "KAC", Name: "Kiaracondong", City: "Bandung"},
	{ID: "YK", Name: "Yogyakarta", City: "Yogyakarta"},
	{ID: "LPN", Name: "Lempuyangan", City: "Yogyakarta"},
	{ID: "SRG", Name: "Semarang Tawang", City: "Semarang"},
	{ID: "SMC", Name: "Semarang Poncol", City: "Semarang"},
	{ID: "SB", Name: "Surabaya Gubeng", City: "Surabaya"},
	{ID: "SBI", Name: "Surabaya Pasar Turi", City: "Surabaya"},
	{ID: "ML", Name: "Malang", City: "Malang"},
	{ID: "MGA", Name: "Malang Kota Lama", City: "Malang"},
	{ID: "SLO", Name: "Solo Balapan", City: "Solo"},
	{ID: "PWS", Name: "Purwosari", City: "Solo"},
	{ID: "CN", Name: "Cirebon", City: "Cirebon"},
	{ID: "CNP", Name: "Cirebon Prujakan", City: "Cirebon"},
	{ID: "TGL", Name: "Tegal", City: "Tegal"},
	{ID: "BOO", Name: "Bogor", City: "Bogor"},
	{ID: "BKS", Name: "Bekasi", City: "Bekasi"},
	{ID: "JGJ", Name: "Jember", City: "Jember"},
}

func getStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stations)
}

// Handler untuk mendapatkan stasiun berdasarkan ID
func getStationByID(w http.ResponseWriter, r *http.Request) {
	id := strings.ToUpper(r.URL.Query().Get("id")) // Ubah ID ke huruf besar
	w.Header().Set("Content-Type", "application/json")

	for _, station := range stations {
		if station.ID == id {
			json.NewEncoder(w).Encode(station)
			return
		}
	}

	// Jika tidak ditemukan
	http.Error(w, `{"error": "Station not found"}`, http.StatusNotFound)
}

// Fungsi utama
func main() {
	// Rute API
	http.HandleFunc("/stations", getStations)
	http.HandleFunc("/station", getStationByID)

	// Menjalankan server di port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
