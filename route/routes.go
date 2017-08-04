package route

import(
    "net/http"

    "service_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        service_handler.Index,
    },
    Route{
        "Login",
        "POST",
        "/login",
        service_handler.Login,
    },
    Route{
        "Add Panitia",
        "POST",
        "/addPanitia",
        service_handler.AddPanitia,
    },
    Route{
        "Add Konsumsi",
        "POST",
        "/addKonsumsi",
        service_handler.AddKonsumsi,
    },
    Route{
        "Add Transport",
        "POST",
        "/addTransport",
        service_handler.AddTransport,
    },
    Route{
        "Add Hotel",
        "POST",
        "/addHotel",
        service_handler.AddHotel,
    },
    Route{
        "Add Kamar to Hotel",
        "POST",
        "/addKamarToHotel",
        service_handler.AddKamarToHotel,
    },
    Route{
        "Register Peserta",
        "POST",
        "/register",
        service_handler.Register,
    },
    Route{
        "Get All Hotel",
        "GET",
        "/getAllHotel",
        service_handler.GetAllHotel,
    },
    Route{
        "Get Hotel By Id",
        "GET",
        "/getHotelById",
        service_handler.GetHotelById,
    },
    Route{
        "Get Room By Hotel Name",
        "GET",
        "/getTipeKamarByName",
        service_handler.GetRoomByName,
    },
    Route{
        "Check Username Exist",
        "GET",
        "/checkUsername",
        service_handler.CheckUsername,
    },
    Route{
        "Get Panitia",
        "GET",
        "/getAllPanitia",
        service_handler.GetPanitia,
    },
    Route{
        "Get Konsumsi",
        "GET",
        "/getAllKonsumsi",
        service_handler.GetKonsumsi,
    },
    Route{
        "Get Transport",
        "GET",
        "/getAllTransport",
        service_handler.GetTransport,
    },
    Route{
        "Get All Peserta",
        "GET",
        "/getAllPeserta",
        service_handler.GetPeserta,
    },
    Route{
        "Search Peserta",
        "GET",
        "/searchPeserta",
        service_handler.SearchPeserta,
    },
    Route{
        "Generate csv",
        "GET",
        "/exportCsv",
        service_handler.ExportCsv,
    },
}