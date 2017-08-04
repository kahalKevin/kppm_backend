package service_handler

import (
    "encoding/json"
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "log"
    "io"
    "io/ioutil"

    "db_handler"
)


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "KPPM")
}

func Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var login db_handler.LoginData
    json.Unmarshal(body, &login)
    hash := db_handler.GetPassword(login.Username)

    if hash == "Not Exist"{
        w.WriteHeader(http.StatusNotFound)
        status := fmt.Sprintf("%d", http.StatusNotFound)
        response := Response{
            Status:     status,
            Info:      "username not exist"}
        json.NewEncoder(w).Encode(response)        
        return
    }

    if strings.Contains(hash, "Lost Connection"){
        w.WriteHeader(http.StatusServiceUnavailable)
        status := fmt.Sprintf("%d", http.StatusServiceUnavailable)
        response := Response{
            Status:     status,
            Info:      "Lost db Connection"}
        json.NewEncoder(w).Encode(response)
        return
    }

    match := CheckPasswordHash(login.Password, hash) 

    if !match{
        w.WriteHeader(http.StatusUnauthorized)
        status := fmt.Sprintf("%d", http.StatusUnauthorized)
        response := Response{
            Status:     status,
            Info:      "Wrong Password"}
        json.NewEncoder(w).Encode(response)
        return
    }

    role := db_handler.GetRole(login.Username)
    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:     status,
        Info:       role}

    json.NewEncoder(w).Encode(response)
}

func AddPanitia(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var panitia db_handler.Panitia
    json.Unmarshal(body, &panitia)

    panitia.Password, _ = HashPassword(panitia.Password)

    id, err := db_handler.CreateCredential(panitia.Username, panitia.Password, "panitia")
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create credential"}
        json.NewEncoder(w).Encode(response)        
        return        
    }

    err = db_handler.CreatePanitia(panitia, id)
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create panitia"}
        json.NewEncoder(w).Encode(response)        
        return        
    }

    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:     status,
        Info:      "success"}

    json.NewEncoder(w).Encode(response)    

}

func AddKonsumsi(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var konsumsi db_handler.Konsumsi
    json.Unmarshal(body, &konsumsi)

    err := db_handler.CreateKonsumsi(konsumsi)
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create konsumsi"}
        json.NewEncoder(w).Encode(response)        
        return        
    }

    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:     status,
        Info:      "success"}

    json.NewEncoder(w).Encode(response)        
}

func AddTransport(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var transport db_handler.Transport
    json.Unmarshal(body, &transport)

    err := db_handler.CreateTransport(transport)
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create transport"}
        json.NewEncoder(w).Encode(response)        
        return        
    }

    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:     status,
        Info:      "success"}

    json.NewEncoder(w).Encode(response)        
}

func AddHotel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var hotel db_handler.Hotel
    json.Unmarshal(body, &hotel)

    id, err := db_handler.CreateHotel(hotel.Nama)
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create hotel"}
        json.NewEncoder(w).Encode(response)        
        return        
    }

    for _, room := range hotel.Kamar{
        db_handler.CreateRoom(room, id)
    }

    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:     status,
        Info:      "success"}

    json.NewEncoder(w).Encode(response)        
}

func AddKamarToHotel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var room db_handler.Room
    json.Unmarshal(body, &room)

    db_handler.CreateRoom(room, room.Idhotel)

    status := fmt.Sprintf("%d", http.StatusAccepted)
    w.WriteHeader(http.StatusAccepted)
 
    response := Response{
        Status:     status,
        Info:      "accepted"}

    json.NewEncoder(w).Encode(response)        
}

func Register(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //restricted 4kb
    body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 4000))
    var peserta db_handler.Peserta
    json.Unmarshal(body, &peserta)

    peserta.Password, _ = HashPassword(peserta.Password) 
    peserta.Kodepeserta = CreateNewKodePeserta(peserta.Role)
    log.Printf("%v", peserta)
    id, err := db_handler.CreateCredential(peserta.Kodepeserta, peserta.Password, peserta.Role)
    if err!=nil{
        w.WriteHeader(http.StatusConflict)
        status := fmt.Sprintf("%d", http.StatusConflict)
        response := Response{
            Status:     status,
            Info:      "fail create kodepeserta"}
        json.NewEncoder(w).Encode(response)        
        return        
    }
    db_handler.CreatePeserta(peserta, id)

    status := fmt.Sprintf("%d", http.StatusAccepted)
    w.WriteHeader(http.StatusAccepted)
 
    response := Response{
        Status:    status,
        Info:      peserta.Kodepeserta}

    json.NewEncoder(w).Encode(response)
}

func GetAllHotel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    hotels := db_handler.GetHotels()

    var _hotels db_handler.Hotelst
    _hotels.Content = hotels

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_hotels)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    idhotel := r.URL.Query().Get("id")
    _idhotel, _ := strconv.Atoi(idhotel)

    hotel := db_handler.GetHotel(_idhotel)

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(hotel)
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    username := r.URL.Query().Get("username")

    info := db_handler.CheckUsernameExist(username)

    status := fmt.Sprintf("%d", http.StatusOK)
    w.WriteHeader(http.StatusOK)
 
    response := Response{
        Status:    status,
        Info:      info}

    json.NewEncoder(w).Encode(response)
}

func GetPanitia(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    panitia := db_handler.GetAllPanitia()

    var _panitia db_handler.Panitiast
    _panitia.Content = panitia

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_panitia)
}

func GetKonsumsi(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    konsumsi := db_handler.GetAllKonsumsi()

    var _konsumsi db_handler.Konsumsist
    _konsumsi.Content = konsumsi

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_konsumsi)
}

func GetTransport(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    trip := r.URL.Query().Get("trip")

    transport := db_handler.GetAllTransport(trip)

    var _transport db_handler.Transportst
    _transport.Content = transport

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_transport)
}

func GetPeserta(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    point := r.URL.Query().Get("idpoint")
    direction := r.URL.Query().Get("direction")

    pesertas := db_handler.GetAllPeserta(point, direction)

    var _pesertas db_handler.Pesertast
    _pesertas.Content = pesertas

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_pesertas)
}

func SearchPeserta(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    query := r.URL.Query().Get("query")

    peserta := db_handler.SearchPeserta(query)

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(peserta)
}

func GetRoomByName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    hotel := r.URL.Query().Get("hotel")

    hotelid := db_handler.GetHotelIdByName(hotel)
    rooms := db_handler.GetRoom(hotelid)

    var _rooms db_handler.Roomst
    _rooms.Content = rooms

    w.WriteHeader(http.StatusOK)
 
    json.NewEncoder(w).Encode(_rooms)
}

func ExportCsv(w http.ResponseWriter, r *http.Request) {
    err := db_handler.ExportCsv()
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(err)
        return
    }
    data, err := ioutil.ReadFile(string("/var/lib/mysql/kppm_dev/data_peserta_kppm.csv"))
    if err == nil {
        w.Header().Set("Content-Disposition", "attachment; filename=data_peserta_kppm.csv")
        w.Header().Set("Content-Type", "text/csv")
        w.WriteHeader(http.StatusOK)
        w.Write(data)
    }else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(err)
        return
    }    
}

