package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/superp00t/etc/yo"
	"github.com/superp00t/gophercraft/datapack"
	"github.com/superp00t/gophercraft/datapack/csv"
	"github.com/superp00t/gophercraft/format/dbc"
	"github.com/superp00t/gophercraft/gcore"
)

func getDBC(name string) *dbc.DBC {
	bts, err := pool.ReadFile(name)
	if err != nil {
		yo.Fatal(name, err)
	}

	dd, err := dbc.Parse(version, bts)
	if err != nil {
		yo.Fatal(name, err)
	}

	return dd
}

func extractDBC(path, to string, sli interface{}) {
	path = "DBFilesClient\\" + path + ".dbc" // TODO: replace extension depending on version

	db := getDBC(path)

	if err := db.ParseRecords(sli); err != nil {
		fmt.Println(path)
		panic(err)
	}

	slice := reflect.TypeOf(sli).Elem()

	switch slice.Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(sli).Elem()

		file, err := pack.WriteFile(to)
		if err != nil {
			panic(err)
		}

		cFile := csv.NewWriter(file)

		for i := 0; i < s.Len(); i++ {
			if err := cFile.AddRecord(s.Index(i).Interface()); err != nil {
				panic(err)
			}
		}

		cFile.Close()
	}
}

func copyFile(in string, out string) {
	b, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}

	pack.WriteBytes(out, b)
}

func generateDatapack(gamePath, dPack string) {
	var err error
	// Open pack folder in temp
	pack, err = datapack.Author(datapack.PackConfig{
		Name:        "Base",
		Description: "Gophercraft Base Content Pack",
		Author:      "DO NOT EDIT: generated by gcraft_gen_datapack on " + time.Now().String(),
		Version:     gcore.Version,
		Depends: []string{
			"@Gophercraft>=" + gcore.Version,
			"@GameVersion==" + fmt.Sprintf("%d", uint32(version))},
	})

	if err != nil {
		yo.Fatal(err)
	}

	var cso []dbc.Ent_CharStartOutfit
	extractDBC("CharStartOutfit", "DB/DBC_CharStartOutfit.csv", &cso)

	var emotesText []dbc.Ent_EmotesText
	extractDBC("EmotesText", "DB/DBC_EmotesText.csv", &emotesText)

	var emotes []dbc.Ent_Emotes
	extractDBC("Emotes", "DB/DBC_Emotes.csv", &emotes)

	var races []dbc.Ent_ChrRaces
	extractDBC("ChrRaces", "DB/DBC_ChrRaces.csv", &races)

	var classes []dbc.Ent_ChrClasses
	extractDBC("ChrClasses", "DB/DBC_ChrClasses.csv", &classes)

	var areaTrigger []dbc.Ent_AreaTrigger
	extractDBC("AreaTrigger", "DB/DBC_AreaTrigger.csv", &areaTrigger)

	// var emotesTextData []dbc.Ent_EmotesTextData
	// extractDBC("EmotesTextData", "DB/DBC_EmotesTextData.csv", &emotesTextData)

	fmt.Println("Generating", dPack, "...")

	if err := pack.ZipToFile(dPack); err != nil {
		panic(err)
	}

	pack.Delete()
}
