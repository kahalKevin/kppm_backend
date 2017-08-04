package db_handler

import (
	"fmt"
	"log"
	"time"
	"os"
	"strconv"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

var Db *sql.DB

func CreateCredential(username, password, role string) (int, error){	
	err := Db.Ping()
	if err != nil {
		return -1, err
	}

	query_insert := "insert into credential(`username`, `password`, `role`) values ('"+username+"', '"+password+"' ,'"+role+"')"

	_, err = Db.Exec(query_insert)
	if err != nil {
		return -1, err
	}

	credId := -1
	Db.QueryRow("select id from credential where username = '"+username+"'").Scan(&credId)

	return credId, nil
}

func CreatePanitia(panitia Panitia, cred_id int) error{
	err := Db.Ping()
	if err != nil {
		return err
	}

	query_insert := "insert into committee(`name`, `ktp`, `phone`, `cred_id`) values ('"+panitia.Name+"', '"+panitia.Ktp+"' ,'"+panitia.Phone+"' ,'"+strconv.Itoa(cred_id)+"')"
	fmt.Println(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		return err
	}
	return nil
}

func CreateKonsumsi(konsumsi Konsumsi) error{
	err := Db.Ping()
	if err != nil {
		return err
	}	
	var harga_str string
	harga_str = strconv.FormatFloat(konsumsi.Harga, 'f', 2, 64)
	
	query_insert := "insert into consumption(`period`, `price`) values ('"+strconv.Itoa(konsumsi.Period)+"', '"+harga_str+"')"
	fmt.Println(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		return err
	}
	return nil
}

func CreateTransport(transport Transport) error{
	err := Db.Ping()
	if err != nil {
		return err
	}	
	var harga_str string
	harga_str = strconv.FormatFloat(transport.Harga, 'f', 2, 64)
	
	query_insert := "insert into transport(`trip`, `type`, `price`) values ('"+transport.Trip+"', '"+transport.Jenis+"', '"+harga_str+"')"
	fmt.Println(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		return err
	}
	return nil
}

func CreateHotel(hotelname string) (int, error){	
	err := Db.Ping()
	if err != nil {
		return -1, err
	}

	query_insert := "insert into hotel(`hotel_name`) values ('"+hotelname+"')"

	res, err := Db.Exec(query_insert)
	if err != nil {
		return -1, err
	}

	hotelId, _ := res.LastInsertId()

	return int(hotelId), nil
}

func CreateRoom(room Room, hotel_id int){
	err := Db.Ping()
	if err != nil {
		fmt.Println("%v", err)
	}

	var single_str string
	single_str = strconv.FormatFloat(room.Single, 'f', 2, 64)
	// double_str = strconv.FormatFloat(room.Double, 'f', 2, 64)
	// ekstrabed_str = strconv.FormatFloat(room.Extrabed, 'f', 2, 64)

	query_insert := "insert into room(`name`, `single`, `hotel_id`) values ('"+room.Nama+"', '"+single_str+"', '"+strconv.Itoa(hotel_id)+"')"
	fmt.Println(query_insert)

	_, err = Db.Exec(query_insert)
	if err != nil {
		fmt.Println("%v", err)
	}
}

func CreatePeserta(peserta Peserta, cred_id int){	
	err := Db.Ping()
	if err != nil {
		log.Println("%v", err)
		return
	}
	_withwife, _cons1, _cons2, _urus1, _urus2 := "0", "0", "0", "0", "0"
	if peserta.WithWife {
		_withwife = "1"
	}
	if peserta.Konsumsi1 {
		_cons1 = "1"
	}
	if peserta.Konsumsi2 {
		_cons2 = "1"
	}
	if peserta.BandaraHotel.Diurus {
		_urus1 = "1"
	}
	if peserta.HotelAcara.Diurus {
		_urus2 = "1"
	}

	var harga_str string
	harga_str = strconv.FormatFloat(peserta.TotalHarga, 'f', 2, 64)

	query_insert := `insert into participant(`+"`"+`cred_id`+"`, "+
											   "`"+`nama`+"`, "+
											   "`"+`jabatan`+"`, "+
											   "`"+`gereja`+"`, "+				
											   "`"+`ktp`+"`, "+
											   "`"+`phone`+"`, "+
											   "`"+`age`+"`, "+
											   "`"+`with_wife`+"`, "+
											   "`"+`wife_name`+"`, "+
											   "`"+`provinsi`+"`, "+
											   "`"+`kota`+"`, "+
											   "`"+`alamat`+"`, "+
											   "`"+`b_h_diurus`+"`, "+
											   "`"+`b_h_jenis`+"`, "+
											   "`"+`b_h_datang`+"`, "+
											   "`"+`b_h_pulang`+"`, "+
											   "`"+`b_h_keterangan`+"`, "+
											   "`"+`h_a_diurus`+"`, "+
											   "`"+`h_a_jenis`+"`, "+
											   "`"+`h_a_keterangan`+"`, "+
											   "`"+`hotel`+"`, "+
											   "`"+`room`+"`, "+
											   "`"+`kasur`+"`, "+
											   "`"+`konsumsi1`+"`, "+
											   "`"+`konsumsi2`+"`, "+
											   "`"+`total_price`+"`, "+
											   "`"+`code`+"`, "+
											   "`"+`created_at`+"`"+
											`) 

			       	 values ('`+strconv.Itoa(cred_id)+`',
			       	 		 '`+peserta.Nama+`',
			       	 		 '`+peserta.Jabatan+`',
			       	 		 '`+peserta.Gereja+`',
			       	 		 '`+peserta.Ktp+`',
			       	 		 '`+peserta.Phone+`',
			       	 		 '`+strconv.Itoa(peserta.Umur)+`',
			       	 		 '`+_withwife+`',
			       	 		 '`+peserta.NameWife+`',
			       	 		 '`+peserta.Provinsi+`',
			       	 		 '`+peserta.Kota+`',
			       	 		 '`+peserta.Alamat+`',
			       	 		 '`+_urus1+`',
			       	 		 '`+peserta.BandaraHotel.Jenis+`',
			       	 		 '`+peserta.BandaraHotel.Waktudatang+`',
			       	 		 '`+peserta.BandaraHotel.Waktupulang+`',
			       	 		 '`+peserta.BandaraHotel.Keterangan+`',
			       	 		 '`+_urus2+`',
			       	 		 '`+peserta.HotelAcara.Jenis+`',
			       	 		 '`+peserta.HotelAcara.Keterangan+`',
			       	 		 '`+peserta.Hotel+`',
			       	 		 '`+peserta.TipeKamar+`',
			       	 		 '`+strconv.Itoa(peserta.Kasur)+`',
			       	 		 '`+_cons1+`',
			       	 		 '`+_cons2+`',
			       	 		 '`+harga_str+`',
			       	 		 '`+peserta.Kodepeserta+`',
			       	 		 '`+time.Now().Format("2006-01-02 15:04:05")+`'
			       	 )` 
			       	 // TIME FORMAT HAS TO BE ALWAYS REFERING TO Mon Jan 2 15:04:05 -0700 MST 2006

	_, err = Db.Exec(query_insert)
	log.Println(query_insert)
	if err != nil {
		log.Println("%v", err)
		return
	}
}

func GetPassword(username string) string{
	err := Db.Ping()
	if err != nil {
		return fmt.Sprintf("Lost Connection %s", err)
	}

	password := "Not Exist"
	err = Db.QueryRow("select password from credential where username = '"+username+"'").Scan(&password)
	if err != nil {
		log.Println(err)
	}
	return password
}

func GetRole(username string) string{
	err := Db.Ping()
	if err != nil {
		return fmt.Sprintf("Lost Connection %s", err)
	}

	role := "none"
	err = Db.QueryRow("select role from credential where username = '"+username+"'").Scan(&role)
	if err != nil {
		log.Println(err)
	}
	return role
}

func GetCredential(id int) (string, string){
	err := Db.Ping()
	if err != nil {
		return "",""
	}

	username, password := "", ""
	err = Db.QueryRow("select username, password from credential where id = '"+strconv.Itoa(id)+"'").Scan(&username,&password)
	if err != nil {
		log.Println(err)
	}
	return username, password
}

func CheckRoleExist(role string) bool{
	err := Db.Ping()
	if err != nil {
		return false
	}

	role_exist := false
	err = Db.QueryRow("select exists(select id from credential where role = '"+role+"' limit 1) as 'exist'").Scan(&role_exist)
	if err != nil {
		log.Println(err)
	}
	return role_exist
}

func CheckUsernameExist(username string) string{
	username_status := "unavailable"
	err := Db.Ping()
	if err != nil {
		return username_status
	}

	exists := true
	err = Db.QueryRow("select exists(select id from credential where username = '"+username+"' limit 1) as 'exist'").Scan(&exists)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		username_status = "available"
	}
	return username_status
}

func GetNumberRegistered(role string) int{
	err := Db.Ping()
	if err != nil {
		return 0
	}

	number_registered := 0
	err = Db.QueryRow("SELECT count(id) FROM credential where role = '"+role+"'").Scan(&number_registered)
	if err != nil {
		log.Println(err)
	}
	return number_registered
}

func GetHotels() Hotels{
	err := Db.Ping()
	if err != nil {
		return nil
	}

	rows, err := Db.Query("select id, hotel_name from hotel")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Hotels Hotels
	var hotel1 Hotel
	for rows.Next() {
		err := rows.Scan(&hotel1.Id, &hotel1.Nama)
		if err != nil {
			log.Println(err)
		}
		hotel1.Kamar = GetRoom(hotel1.Id)
		hotel1.Jumlah = len(hotel1.Kamar)
		_Hotels = append(_Hotels, hotel1)
	} 
	return _Hotels
}

func GetHotel(idhotel int) Hotel{
	var hotel1 Hotel

	err := Db.Ping()
	if err != nil {
		return hotel1
	}

	err = Db.QueryRow("select id, hotel_name from hotel where id =" + strconv.Itoa(idhotel)).Scan(&hotel1.Id, &hotel1.Nama)
	if err != nil {
		log.Println(err)
	}
	hotel1.Kamar = GetRoom(hotel1.Id)
	hotel1.Jumlah = len(hotel1.Kamar)
	
	return hotel1
}

func GetHotelIdByName(hotelname string) int{
	var hotelid int

	err := Db.Ping()
	if err != nil {
		return -1
	}

	err = Db.QueryRow("select id from hotel where hotel_name = '"+hotelname+"'").Scan(&hotelid)
	if err != nil {
		return -1
	}
	
	return hotelid
}

func GetRoom(idhotel int) Rooms{
	err := Db.Ping()
	if err != nil {
		return nil
	}

	rows, err := Db.Query("select id, name, single from room where hotel_id =" + strconv.Itoa(idhotel))
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Rooms Rooms
	var room1 Room
	for rows.Next() {
		err := rows.Scan(&room1.Id, &room1.Nama, &room1.Single)
		if err != nil {
			log.Println(err)
		}
		_Rooms = append(_Rooms, room1)
	} 
	return _Rooms
}

func GetAllPanitia() Panitias{
	err := Db.Ping()
	if err != nil {
		return nil
	}
	var cred_id int
	rows, err := Db.Query("select id, name, ktp, phone, cred_id from committee")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Panitias Panitias
	var panitia1 Panitia
	for rows.Next() {
		err := rows.Scan(&panitia1.Id, &panitia1.Name, &panitia1.Ktp, &panitia1.Phone, &cred_id)
		if err != nil {
			log.Println(err)
		}
		panitia1.Username, panitia1.Password = GetCredential(cred_id)
		_Panitias = append(_Panitias, panitia1)
	} 
	return _Panitias
}

func GetAllKonsumsi() Konsumsis{
	err := Db.Ping()
	if err != nil {
		return nil
	}
	rows, err := Db.Query("select id, period, price from consumption")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Konsumsis Konsumsis
	var konsumsi1 Konsumsi
	for rows.Next() {
		err := rows.Scan(&konsumsi1.Id, &konsumsi1.Period, &konsumsi1.Harga)
		if err != nil {
			log.Println(err)
		}
		_Konsumsis = append(_Konsumsis, konsumsi1)
	} 
	return _Konsumsis
}

func GetAllTransport(trip string) Transports{
	err := Db.Ping()
	if err != nil {
		return nil
	}
	trip_type := "'bandara-hotel'"
	if trip == "2" {
		trip_type = "'hotel-acara'"
	}
	rows, err := Db.Query("select id, trip, type, price from transport where trip = "+trip_type)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Transports Transports
	var transport1 Transport
	for rows.Next() {
		err := rows.Scan(&transport1.Id, &transport1.Trip, &transport1.Jenis, &transport1.Harga)
		if err != nil {
			log.Println(err)
		}
		_Transports = append(_Transports, transport1)
	} 
	return _Transports
}

func GetAllPeserta(point, direction string) Pesertas{
	err := Db.Ping()
	if err != nil {
		return nil
	}
	var query_peserta string
	if direction == "next" {
		query_peserta = `SELECT participant.id, code, nama, role FROM participant
							JOIN credential
							ON credential.id = participant.cred_id
						WHERE participant.id>`+point+`
						ORDER BY participant.id ASC
						LIMIT 50`
	}else{
		query_peserta = `SELECT id, code, nama, role FROM(
							SELECT participant.id, code, nama, role FROM participant
								JOIN credential
								ON credential.id = participant.cred_id
							WHERE participant.id<`+point+`
							ORDER BY participant.id DESC
							LIMIT 50
						) AS Table1
						ORDER BY id ASC`
	}

	rows, err := Db.Query(query_peserta)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var _Pesertas Pesertas
	var peserta1 Peserta
	for rows.Next() {
		err := rows.Scan(&peserta1.Id, &peserta1.Kodepeserta, &peserta1.Nama, &peserta1.Role)
		if err != nil {
			log.Println(err)
		}
		_Pesertas = append(_Pesertas, peserta1)
	} 
	return _Pesertas
}

func SearchPeserta(keyword string) Peserta{	
	var peserta1 Peserta
	err := Db.Ping()
	if err != nil {
		return peserta1
	}

	query_select := `SELECT 
						code, nama, role, jabatan, gereja,
						ktp, phone, age, with_wife, wife_name,
					    provinsi, kota, alamat, b_h_diurus,
					    b_h_jenis, b_h_datang, b_h_pulang,
					    b_h_keterangan, h_a_diurus,
					    h_a_jenis, h_a_keterangan,
					    hotel, room, kasur,
					    konsumsi1, konsumsi2,
					    total_price
					FROM participant
					JOIN credential
					ON credential.id = participant.cred_id
					WHERE (
						code = '`+keyword+`'
					    OR
					    nama = '`+keyword+`'
					)
					LIMIT 1`

	Db.QueryRow(query_select).Scan(&peserta1.Kodepeserta, &peserta1.Nama, &peserta1.Role,
								   &peserta1.Jabatan, &peserta1.Gereja, &peserta1.Ktp,
								   &peserta1.Phone, &peserta1.Umur, &peserta1.WithWife,
								   &peserta1.NameWife, &peserta1.Provinsi, &peserta1.Kota,
								   &peserta1.Alamat, &peserta1.BandaraHotel.Diurus,
								   &peserta1.BandaraHotel.Jenis, &peserta1.BandaraHotel.Waktudatang,
								   &peserta1.BandaraHotel.Waktupulang, &peserta1.BandaraHotel.Keterangan,
								   &peserta1.HotelAcara.Diurus, &peserta1.HotelAcara.Jenis,
								   &peserta1.HotelAcara.Keterangan, &peserta1.Hotel, &peserta1.TipeKamar,
								   &peserta1.Kasur, &peserta1.Konsumsi1, &peserta1.Konsumsi2,
								   &peserta1.TotalHarga)

	return peserta1
}


func ExportCsv() error{	
	err := os.Remove("/var/lib/mysql/kppm_dev/data_peserta_kppm.csv")
	if err != nil {
		log.Println("data peserta not exist")
	}
	query_export := `SELECT 'KodePeserta', 'Nama', 'Jenis', 'Jabatan', 'OrganisasiGereja',
					   'KTP', 'Handphone', 'Umur', 'DenganIstri', 'NamaIstri',
				       'Provinsi', 'Kota', 'Alamat', 'Bandara-HotelDiurusPanitia',
					   'TransportasiBandara-Hotel', 'WaktuKedatangan', 'WaktuKepulangan',
					   'KeteranganTransportBandara-Hotel', 'Hotel-AcaraDiurusPanitia',
					   'TransportasiHotel-Acara', 'KeteranganTransportHotel-acara',
					   'Hotel', 'TipeKamar', 'SingleOrDoubleBed',
					   'MemesanKonsumsi', 'MemesanSnack',
					   'TotalHarga', 'WaktuMendaftar'
					UNION ALL
					SELECT 
						code, nama, role, jabatan, gereja,
						ktp, phone, age, with_wife, wife_name,
					    provinsi, kota, alamat, b_h_diurus,
					    b_h_jenis, b_h_datang, b_h_pulang,
					    b_h_keterangan, h_a_diurus,
					    h_a_jenis, h_a_keterangan,
					    hotel, room, kasur,
					    konsumsi1, konsumsi2,
					    total_price, created_at
					FROM participant
					JOIN credential
					ON credential.id = participant.cred_id	INTO OUTFILE 'data_peserta_kppm.csv' 
				FIELDS ENCLOSED BY '"' 
				TERMINATED BY ';' 
				ESCAPED BY '"' 
				LINES TERMINATED BY '\r\n';`

	_, err = Db.Exec(query_export)
	if err != nil {
		return err
	}
	return nil
}

