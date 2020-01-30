package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("small1.txt")
	check(err)
	defer f.Close()

	fi, err := f.Stat()
	check(err)

	var textSize int64 = fi.Size()

	for textSize < 1000000 {
		f.WriteString("The problem of university written exam timetables is a well-known multidimensional constraint matching problem that focuses on assigning courses to faculty members in classes over a limited period of time. Therefore, it is a difficult time-consuming problem that universities face. This complexity is justified by the size of exam planning and the large number of restrictions and allocation criteria. Usually a long and hard job of providing an appropriate and optimized solution. We understand that it is imperative that the need for an automated information system be developed. This paper focuses on the development of an algorithm to automatically satisfy all constraints on the particular problem of the university timetable of the written examinations of the Department of Engineering Information and Communication Systems of the University. Emphasize that the algorithm covers most constraints for the needs of the particular university and can be considered insufficient for some other part, which is perfectly reasonable from the essential differences between the parts.")

		fi, err := f.Stat()
		check(err)

		textSize = fi.Size()

		fmt.Printf("The file is %d bytes long\n", fi.Size())
	}
}
