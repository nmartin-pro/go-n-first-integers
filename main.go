package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		nParam, ok := q["n"]
		if !ok && len(nParam) < 1 {
			http.Error(w, "missing 'n' query param", http.StatusBadRequest)
			return
		}

		n, err := strconv.Atoi(nParam[0])
		if err != nil {
			http.Error(w, "invalid 'n' query param", http.StatusBadRequest)
			return
		}

		res := (n * (n + 1)) / 2

		fmt.Fprintf(w, "Sum of integerstest from 1 to %d is: %d", n, res)
	})

	http.ListenAndServe(":5000", nil)
}
