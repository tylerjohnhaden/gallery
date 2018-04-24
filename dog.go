package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/doggo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%[1]s%[3]c%[1]s%[3]c, %[3]c%[2]s%[3]c, 0x60, 0x22)\n\t})\n\thttp.ListenAndServe(%[4]c:80%[4]c, nil)\n}\n", `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/doggo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%[1]s%[3]c%[1]s%[3]c, %[3]c%[2]s%[3]c, 0x60, 0x22)\n\t})\n\thttp.ListenAndServe(%[4]c:80%[4]c, nil)\n}\n", `, `



                                    oooo\\
                                mmmoo * o\\
                                 wwoooooooo                    //
                                      oooooooooooooooooooooooo/
                                       ooooooooooooooooooooooo
                                      mmm   mmm    mmm   mmm

                                                                               doggo (adv.) Remain motionless and quiet to escape detection.

`, 0x60, 0x22)
	})
	http.ListenAndServe(":80", nil)
}
