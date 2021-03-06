package main
import (
  "bufio"
  "encoding/csv"
  "fmt"
  "bytes"
  "io"
  "io/ioutil"
  "log"
  "os"
  "net/url"
  "strconv"
)

func main() {
  csvFile, _ := os.Open("guids.csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  var urls []string
  u, err := url.Parse("http://hasctd1/members/GetPriPreview.aspx?Width=2550&PriGuid=e8ec7a41-228c-4df4-9ddc-af4c2963263f&ext=.png")
  if err != nil {
    log.Fatal(err)
  }
  for {
    line, error := reader.Read()
    if error == io.EOF {
      break
    } else if error != nil {
      log.Fatal(error)
    }
    q := u.Query()
    q.Set("PriGuid", line[0])
    u.RawQuery = q.Encode()
    urls = append(urls,u.String())
  }
  fmt.Printf("%v",urls)

  var replaceText bytes.Buffer
  for i, u := range urls {
    replaceText.WriteString(fmt.Sprintf("<tr><td><a href='%s'>Link %d</a></td></tr>",u,i))
    if i % 10000 == 0 {
      input, err := ioutil.ReadFile("template.txt")
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      output := bytes.Replace(input, []byte("##CONTENT##"), replaceText.Bytes(), -1)

      if err = ioutil.WriteFile("PriLinks"+strconv.Itoa(i)+".aspx", output, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      replaceText.Reset()
    }
  }

  input, err := ioutil.ReadFile("template.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  output := bytes.Replace(input, []byte("##CONTENT##"), replaceText.Bytes(), -1)

  if err = ioutil.WriteFile("PriLinks.aspx", output, 0666); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
