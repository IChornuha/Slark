package forum

import "golang.org/x/text/encoding/charmap"

//ToUTF decode from Win1251
func ToUTF(post string) string {
	dec := charmap.Windows1251.NewDecoder()
	// newBody := make([]byte, len(post)*2)
	postAsByte := []byte(post)
	// n, _, err := dec.Transform(newBody, postAsByte, true)
	n, err := dec.Bytes(postAsByte)
	if err != nil {
		panic(err)
	}
	newBody := string(n) //newBody[:n]
	return newBody
}
