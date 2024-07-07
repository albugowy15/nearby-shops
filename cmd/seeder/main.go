package main

import (
	"log"

	"github.com/albugowy15/nearby-shops/internal/config"
)

var seed = `
    INSERT INTO shops (name, description, city, location) VALUES ('Toko Pak Tretan', 'Toko Madura Pak tretan buka 24 jam', 'Surabaya', 'SRID=4326;POINT(112.768845 -7.250445)');
    INSERT INTO shops (name, description, city, location) VALUES ('Alfamart Keputih', 'Alfamart Cabang Keputih Surabaya', 'Surabaya', 'SRID=4326;POINT(113.76512 -5.65121)');
    INSERT INTO shops (name, description, city, location) VALUES ('Indomaret Gebang', 'Indomaret cabang gebang Surabaya','Surabaya', 'SRID=4326;POINT(110.12131 -6.773232)');
    INSERT INTO shops (name, description, city, location) VALUES ('Sakinah Supermarket', 'Supermarket sakinah, sedia peralatan alat tulis', 'Surabaya', 'SRID=4326;POINT(112.51281 -7.737231)');
    INSERT INTO shops (name, description, city, location) VALUES ('Toko Bu Atmirah', 'Toko sedia alat tulis keperluan sekolah', 'Surabaya', 'SRID=4326;POINT(113.612121 -5.723223)');
    INSERT INTO shops (name, description, city, location) VALUES ('Skynet Gadget', 'Toko gadget sedia laptop, smartphone, dan printer','Surabaya', 'SRID=4326;POINT(110.83121 -6.73711)');
  `

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	db.MustExec(seed)
	db.Close()
	log.Println("DB Seeder success")
}
