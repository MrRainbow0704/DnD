package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/MrRainbow0704/DnD/internal/log"
	"github.com/MrRainbow0704/DnD/internal/version"
)

type Config struct {
	Host          string        `json:"host"`            // L'indirizzo IP dell'host.
	Port          string        `json:"port"`            // La porta usata dal server.
	Address       string        `json:"-"`               // Combinazione di [config.Host] e [config.Port], separati da due punti.
	DBPath        string        `json:"dbPath"`          // Il percorso del database SQLite.
	PasswdPepper  string        `json:"passwordPepper"`  // Il pepe per l'hashing delle password.
	PasswdTime    uint32        `json:"passwordTime"`    // Il tempo per l'hashing delle password.
	PasswdMemory  uint32        `json:"passwordMemory"`  // La memoria per l'hashing delle password.
	PasswdThreads uint8         `json:"passwordThreads"` // I thread per l'hashing delle password.
	PasswdKeyLen  uint32        `json:"passwordKeyLen"`  // La lunghezza della chiave per l'hashing delle password.
	PasswdSaltLen uint32        `json:"passwordSaltLen"` // La lunghezza del sale per l'hashing delle password.
	JWTKey        string        `json:"jwtKey"`          // La chiave per i JSON Web Token.
	JWTMaxAge     int           `json:"jwtMaxAge"`       // LDuranta in secondi per cui il JWT è valido.
	CacheTime     time.Duration `json:"cacheTime"`       // Il tempo per cui salvare la cache.
}

var (
	cnf     Config
	cnfPath string
)

func init() {
	if version.IsDev() {
		cnfPath = "./configs/dev.json"
	} else {
		cnfPath = "./configs/prod.json"
	}
	data, err := os.ReadFile(cnfPath)
	if err != nil {
		log.Panicf("Errore durante il caricamento della configurazione: %s", map[string]any{"file": cnfPath}, err)
	}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		log.Panicf("Errore durante la decompilazione della configurazione: %s", map[string]any{"file": cnfPath}, err)
	}
	cnf.Address = cnf.Host + ":" + cnf.Port
	cnf.CacheTime = cnf.CacheTime * 1_000_000_000 // Converte da nanosecondi a secondi
}

func Get() *Config {
	return &cnf
}
