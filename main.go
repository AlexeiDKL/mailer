package main

import (
	"fmt"
	"log/slog"

	"net/http"
	"os"

	"dkl.dklsa.mailer/iternal/config"
	"dkl.dklsa.mailer/iternal/handlers"
	"dkl.dklsa.mailer/iternal/middleware"
	"dkl.dklsa.mailer/iternal/storage"
	sqlites "dkl.dklsa.mailer/iternal/storage/sqlite"

	"github.com/gorilla/mux"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

/*
$Env:GOOS="windows"
*/

func initAll() {
	cfg := initConfig()

	log := setupLogger(cfg.Env)

	log.Debug("config", cfg)

	initDB()

	initServer(cfg.HTTPServer.Address)
}

func main() {
	stor, err := sqlites.CreateCompanyTable(`storage\storage.db`)
	q := sqlites.CreateCompanyStorages(stor)

	t, err := q.Insert("dkl")
	fmt.Println(t)
	fmt.Println(err)

	i, err := q.Select(storage.Pair{Type: "names", Value: "dkl"})
	fmt.Println(err)
	fmt.Println(i.ID)
	fmt.Println(i.Name)

	// initAll()
}

func initServer(url string) {
	router := mux.NewRouter()

	handlers := [](handlers.Handlers){
		handlers.NewUrlHandler(),
		handlers.NewUrlTypeHandler(),
		handlers.NewDomainHandler(),
		handlers.NewCompanyHandler(),
	}

	for _, h := range handlers { //todo
		router.HandleFunc(h.UrlCreate(), h.Create)
		router.HandleFunc(h.UrlGet(), h.Get).Methods("GET")
		router.HandleFunc(h.UrlEdit(), h.Edit).Methods("PUT")
		router.HandleFunc(h.UrlDelete(), h.Delete).Methods("DELETE")
	}

	http.Handle("/", router)
	handler := middleware.SetMiddleware(router)

	slog.Info(url)
	http.ListenAndServe(url, handler)
}

func initConfig() *config.Config {
	os.Setenv("CONFIG_PATH", "./config/local.yaml")
	return config.MustLoad()
}

func initDB() {
}

// setupLogger - устанавливает логгер в зависимости от среды
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	// log = log.With(slog.String("env", env))

	return log
}

/*
	Заполнение бд /createcompany
		компании
		домены
		ссылки

	Регистрация устройства
		создание пинкода /getpin (получаем мыло)
			--получаем домен, и смотрим, есть ли такой в таблице "domens"
			-- если нет, то отправляем ошибку
			--если домен есть, то
			    --генерируем пин
				--записываем в таблицу
				--записываем в мейлер
				--возвращаем пин
		валидация пинкода /validpin
			--сравниваем пинкоды
			--если пинкоды не совпадают, то отправляем ошибку
			--если пинкоды совпадают, то
				--получаем список серверов по домену мыла

*/

/*
	иницилизация бд
		если надо, создаём таблицы
		если не надо, иницилизируем
	Иницилизация мейлера
		получаем данные мейлера из конфигурационного файлаып
		    если нет конфигурационного файла, то отправляем ошибку
	Мейлер
		Проверяем данные в бд (есть ли новые письма)
		Если есть новые письма, отправляем их

	api
		/getpin
			получаем маил
			если
				домена нет, то отправляем ошибку
			иначе
				генерируем пин
				сохраняем пин в бд
				отправляем пин на мейл
		/getservers
			получаем пин и маил
			сравниваем пин в бд по маил
			если пин верный, удаляем его из бд
				отправляем список серверов связаных с доменом емаила
		    если пин неверный, отправляем ошибку
*/
