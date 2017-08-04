package db_handler

const (
Single = iota
Double
SingleExtra
DoubleExtra
)

type Response struct {
    Status 	    string    `json:"status"`
    Info	 	string    `json:"info"`
}

type LoginData struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
}

type Panitia struct{
	Id          int       `json:"id"`
    Name        string    `json:"name"`
    Username    string    `json:"username"`
    Password    string    `json:"password"`
    Ktp         string    `json:"ktp"`
    Phone       string    `json:"nohp"`
}
type Panitias []Panitia
type Panitiast struct {
	Content     Panitias    `json:"panitia"`
}

type Konsumsi struct{
	Id          int       `json:"id"`
	Period      int       `json:"jenis"`
	Harga       float64   `json:"harga"`
}
type Konsumsis []Konsumsi
type Konsumsist struct {
	Content     Konsumsis    `json:"konsumsi"`
}

type Transport struct{
	Id          int       `json:"id"`
	Trip        string    `json:"trip"`
	Jenis       string    `json:"jenis"`
	Harga       float64   `json:"harga"`
}
type Transports []Transport
type Transportst struct {
	Content     Transports    `json:"transport"`
}

type Room struct{
	Id          int       `json:"id"`
	Nama        string    `json:"nama"`
	Single      float64   `json:"harga"`
	Double      float64   `json:"double"`
	Extrabed    float64   `json:"ekstrabed"`
	Idhotel     int       `json:"idhotel"`
}
type Rooms []Room
type Roomst struct {
	Content     Rooms     `json:"tipekamar"`
}

type Hotel struct{
	Id          int       `json:"id"`
	Nama        string    `json:"namahotel"`
	Jumlah      int       `json:"jlhkamar"`
	Kamar       Rooms     `json:"tipekamar"`
}
type Hotels []Hotel
type Hotelst struct {
	Content     Hotels    `json:"hotel"`
}

type Accomodate struct {
	Diurus          bool       `json:"diurus"`
	Jenis           string     `json:"jenis"`
	Waktudatang     string     `json:"waktudatang"`
	Waktupulang     string     `json:"waktupulang"`
	Keterangan      string     `json:"keterangan"`
}

type Peserta struct{
	Id              int         `json:id,omitempty`
	Username        string      `json:"username"`
	Password        string      `json:"password"`
	Kodepeserta     string      `json:"kodepeserta"`
	Nama            string      `json:"nama"`
	Role            string      `json:"role"`
	Jabatan         string      `json:"jabatan"`
	Gereja          string      `json:"gerejaorg"`
	Ktp             string      `json:"ktp"`
	Phone           string      `json:"nohp"`
	Umur            int         `json:"umur"`
	WithWife        bool        `json:"denganistri"`
	NameWife        string      `json:"namaistri"`
	Provinsi        string      `json:"provinsi"`
	Kota            string      `json:"kota"`
	Alamat          string      `json:"alamat"`
	BandaraHotel    Accomodate  `json:"bandarahotel"`
	HotelAcara      Accomodate  `json:"hotelacara"`
	Hotel           string      `json:"hotel"`
	TipeKamar       string      `json:"tipekamar"`
	Kasur           int         `json:"kasur"`
	Konsumsi1       bool        `json:"konsumsi1"`
	Konsumsi2       bool        `json:"konsumsi2"`
	Konsumsi3       bool        `json:"konsumsi3"`
	TotalHarga      float64     `json:"totalharga"`
}
type Pesertas []Peserta
type Pesertast struct {
	Content     Pesertas    `json:"peserta"`
}
