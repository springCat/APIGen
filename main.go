package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Line struct{
	Name   string
	Option string
	Typ    string
	Parent string
	Desc   string
}

type ClassData struct{
	Lines       []Line
	ClassName   string
	PackagePath string
	ErrCodeLines []string
}


var classTemplate = `package {{ .PackagePath }};

import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@Builder
@Data
/**
* {{range .ErrCodeLines}}
*  {{ . }}{{ end }}
* @author springcat
*/
public class {{ .ClassName }}{
	{{range .Lines}}
	/**
	 * must: {{ .Option }}
	 * {{ .Desc }}
	 *
	*/
	private {{ .Typ }} {{ .Name }};
	{{ end }}
`

var subTemplate = `    @Data
    @NoArgsConstructor
    @Builder
    class {{ .ClassName }}{
	{{range .Lines}}
        /**
        * must: {{ .Option }}
        * {{ .Desc }}
        *
        */
        private {{ .Typ }} {{ .Name }};
	{{ end }}
    }
`

func main()  {

	packagePath := flag.String("p", "", "the package of the java")
	className := flag.String("i", "", "the name of the interface")
	dataPath := flag.String("d", "~", "the path of data to generator the java")
	outPath := flag.String("o", "~", "the path of data to output")

	flag.Parse()

	f, err := os.Open(*dataPath)
	assertNotError(err)

	b := bufio.NewReader(f)

	reqLines := make([]string,0);
	resLines := make([]string,0);
	errCodeLines := make([]string,0);

	//hanlde line
	for {
		data, _, err := b.ReadLine()
		if err != nil {
			break
		}
		line := strings.TrimSpace(string(data))

		if len(line) > 0 && line == "#### **请求字段**" {
			for len(line) > 0 && line != "#### **返回字段**" {
				data, _, err = b.ReadLine()
				line = strings.TrimSpace(string(data))
				if len(line) > 0 {
					reqLines = append(reqLines, line)
				}
			}
		}

		if len(line) > 0 && line == "#### **返回字段**" {
			for len(line) > 0 && line != "#### **错误码**" {
				data, _, err = b.ReadLine()
				line = strings.TrimSpace(string(data))
				if len(line) > 0 {
					resLines = append(resLines, line)
				}
			}
		}

		if len(line) > 0 && line == "#### **错误码**" {
			for len(line) > 0 && line != "#### **接口示例**" {
				data, _, err = b.ReadLine()
				line = strings.TrimSpace(string(data))
				if len(line) > 0 {
					errCodeLines = append(errCodeLines, line)
				}
			}
		}

	}
	reqLines = reqLines[2:]
	resLines = resLines[2:]
	errCodeLines = errCodeLines[2:]

	fmt.Println("packagePath: "+*packagePath)
	fmt.Println("interaceName: "+*className)
	fmt.Println("dataPath: "+*dataPath)


	reqJava, err := os.Create(*outPath + "/" + *className + "Request.java")
	assertNotError(err)

	resJava, err := os.Create(*outPath + "/" + *className + "Response.java")
	assertNotError(err)

	genClass(reqLines, nil, *className+"Request", *packagePath,reqJava)
	genClass(resLines, errCodeLines, *className+"Response",*packagePath,resJava)

	fmt.Println(*outPath + "/" + *className + "Request.java gen success")
	fmt.Println(*outPath + "/" + *className + "Response.java gen success")
}

func genClass(lines []string, errCodeLines []string, className string, packagePath string,wr io.Writer)  {
	buffer := map[Line]int{}

	genBuffer := map[string]Line{}

	for _,line := range lines {
		attributes := strings.Split(line, "|")
		line := Line{
			Name:strings.TrimSpace(attributes[1]),
			Option:strings.TrimSpace(attributes[2]),
			Typ:strings.TrimSpace(attributes[3]),
			Parent:strings.TrimSpace(attributes[4]),
			Desc:strings.TrimSpace(attributes[5]),
		}
		buffer[line] = 1;
	}

	//gen root class
	classTemplateExec := template.Must(template.New("classTemplate").Parse(classTemplate))
	handleLines := make([]Line,0);
	for k := range buffer {
		if k.Parent == "" {
			delete(buffer,k)
			genBuffer[k.Name] = k
			handleLines = append(handleLines, k)
		}

	}
	classData := &ClassData{
		Lines:handleLines,
		ClassName:className,
		PackagePath:packagePath,
	}
	//for error in response
	if errCodeLines != nil {
		classData.ErrCodeLines = errCodeLines;
	}

	classTemplateExec.Execute(wr,classData)

	//gen sub class
	subTemplate := template.Must(template.New("subTemplate").Parse(subTemplate))
	for len(buffer) > 0  {

		genBuffer = genSubClass(subTemplate, buffer,genBuffer,wr)

	}

	io.WriteString(wr,"}")
}

func genSubClass(subTemplate *template.Template, buffer map[Line]int,genBuffer map[string]Line,wr io.Writer) ( map[string]Line ) {
	resultGenBuffer := map[string]Line{}
	for genk, genv := range genBuffer {
		handleLines := make([]Line,0);
		for k := range buffer {
			if k.Parent == genk {
				delete(buffer, k)
				resultGenBuffer[k.Name] = k
				handleLines = append(handleLines, k)
			}
		}

		if len(handleLines) > 0 {
			classData := &ClassData{
				Lines:       handleLines,
				ClassName:   genv.Typ,
			}
			subTemplate.Execute(wr, classData)
		}
	}

	return resultGenBuffer
}

func assertNotError(err error)  {
	if err != nil {
		panic(err)
	}
}
