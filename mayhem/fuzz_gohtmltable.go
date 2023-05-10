package fuzz_gohtmltable

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/nfx/go-htmltable"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                testString, _ := fuzzConsumer.GetString()

                htmltable.NewFromString(testString)
                return 0

            case 1:
                testUrl, _ := fuzzConsumer.GetString()

                htmltable.NewFromURL(testUrl)
                return 0

            case 2:
                var testPage htmltable.Page
                fuzzConsumer.GenerateStruct(&testPage)

                testPage.Len()
                return 0

            case 3:
                var testPage htmltable.Page
                fuzzConsumer.GenerateStruct(&testPage)

                column, _ := fuzzConsumer.GetString()

                testPage.FindWithColumns(column)
                return 0

            // case 4:
            //     var testCellSpan htmltable.cellSpan
            //     fuzzConsumer.GenerateStruct(&testCellSpan)

            //     x, _ := fuzzConsumer.GetInt()
            //     y, _ := fuzzConsumer.GetInt()

            //     testCellSpan.Match(x, y)
            //     return 0

            case 5:
                var testTable htmltable.Table
                fuzzConsumer.GenerateStruct(&testTable)

                testTable.String()
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}