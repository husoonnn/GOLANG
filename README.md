An avid programmer learning a new language to upgrade his skillsets:) 

//JSON Marshalling
p1 := Person(name: "joe", add: "a st.")

barr, err := json.Marshal(p1)

//Marshal() returns JSON representation as []byte - byte array

var p2 Person
err := json.Unmarshal (barr, &p2)
//Unmarshal() converts a JSON []byte into a GO object


// ioutil package (used to read files)
dat, e :=
ioutil.ReadFile("test.txt")

dat = "Hello, World"
err := ioutil.WRiteFile("outfile.txt", dat, 0777)


// os package (File access)
os.Open() opens a file
returns a file descriptor(File)
os.Close() closes a file

Opening and reading
f, err := os.Open("dt.txt") //f is file handle that is called by open
barr := make([]byte, 10)
nb, err := f.read(barr)
f.Close()

- reads and fills barr
- read returns # of bytes read
- may be less than []byte length

os File Create/Write
f, err := os.Create("outfile.txt")

barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WRiteString("Hi")