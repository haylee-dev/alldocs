package func_all

import (
	"encoding/base64"
	"github.com/gookit/color"
	"strings"
)

func CleanTitleForUrl(title string) string {
	titleurl0 := strings.Replace(title, "/", "-", -1)
	titleurl1 := strings.Replace(titleurl0, " ", "-", -1)
	titleurl2 := strings.Replace(titleurl1, "(", "", -1)
	titleurl3 := strings.Replace(titleurl2, ")", "", -1)
	titleurl4 := strings.Replace(titleurl3, ".", "-", -1)
	titleurl5 := strings.Replace(titleurl4, "#", "-", -1)
	titleurl6 := strings.Replace(titleurl5, "/", "-", -1)
	titleurl7 := strings.Replace(titleurl6, "&", "", -1)
	titleurl8 := strings.Replace(titleurl7, ":", "", -1)
	titleurl9 := strings.Replace(titleurl8, ",", "-", -1)
	titleurl10 := strings.Replace(titleurl9, "--", "-", -1)
	titleurl := strings.ToLower(titleurl10)
	return titleurl
}
func CleanCategory(categoryname string) string {
	categoryname0 := strings.Replace(categoryname, ",", "", -1)     // remove commas
	categoryname1 := strings.Replace(categoryname0, "/", "", -1)    // remove forward slash
	categoryname2 := strings.Replace(categoryname1, "&", "and", -1) // remove &
	categoryname3 := strings.Replace(categoryname2, " ", "-", -1)   // add slashes
	categoryname4 := strings.Replace(categoryname3, "--", "-", -1)  // remove double slashes
	categoryname_cl := strings.ToLower(categoryname4)
	return categoryname_cl
}
func CreateColorArrayFromTitle(title string) []string {
	colorstr := ""
	title_lc := strings.ToLower(title)
	red := strings.Contains(title_lc, "red")
	blue := strings.Contains(title_lc, "blue")
	yellow := strings.Contains(title_lc, "yellow")
	black := strings.Contains(title_lc, "black")
	white := strings.Contains(title_lc, "white")
	green := strings.Contains(title_lc, "green")
	orange := strings.Contains(title_lc, "orange")
	grey := strings.Contains(title_lc, "grey")
	silver := strings.Contains(title_lc, "silver")
	brown := strings.Contains(title_lc, "brown")
	beige := strings.Contains(title_lc, "beige")
	purple := strings.Contains(title_lc, "purple")
	pink := strings.Contains(title_lc, "pink")
	gold := strings.Contains(title_lc, "gold")

	if red == true {
		colorstr = colorstr + "red" + " "
	}
	if blue == true {
		colorstr = colorstr + "blue" + " "
	}
	if yellow == true {
		colorstr = colorstr + "yellow" + " "
	}
	if black == true {
		colorstr = colorstr + "black" + " "
	}
	if white == true {
		colorstr = colorstr + "white" + " "
	}
	if green == true {
		colorstr = colorstr + "green" + " "
	}
	if orange == true {
		colorstr = colorstr + "orange" + " "
	}
	if grey == true {
		colorstr = colorstr + "grey" + " "
	}
	if silver == true {
		colorstr = colorstr + "silver" + " "
	}
	if brown == true {
		colorstr = colorstr + "brown" + " "
	}
	if beige == true {
		colorstr = colorstr + "beige" + " "
	}
	if purple == true {
		colorstr = colorstr + "purple" + " "
	}
	if pink == true {
		colorstr = colorstr + "pink" + " "
	}
	if gold == true {
		colorstr = colorstr + "gold" + " "
	}

	colorstr_cl := strings.TrimSpace(colorstr) //remove last space
	color.Red.Println(colorstr_cl)             // Access the first element value
	ColorsArray := strings.Split(colorstr_cl, " ")
	color.Red.Println(ColorsArray)

	return ColorsArray
}
func CreateMainColor(title string) string {
	main_color := ""
	title_lc := strings.ToLower(title)
	red := strings.Contains(title_lc, "red")
	blue := strings.Contains(title_lc, "blue")
	yellow := strings.Contains(title_lc, "yellow")
	black := strings.Contains(title_lc, "black")
	white := strings.Contains(title_lc, "white")
	green := strings.Contains(title_lc, "green")
	orange := strings.Contains(title_lc, "orange")
	grey := strings.Contains(title_lc, "grey")
	silver := strings.Contains(title_lc, "silver")
	brown := strings.Contains(title_lc, "brown")
	beige := strings.Contains(title_lc, "beige")
	purple := strings.Contains(title_lc, "purple")
	pink := strings.Contains(title_lc, "pink")
	gold := strings.Contains(title_lc, "gold")

	if red == true {
		main_color = "red"
	}
	if blue == true {
		main_color = "blue"
	}
	if yellow == true {
		main_color = "yellow"
	}
	if black == true {
		main_color = "black"
	}
	if white == true {
		main_color = "white"
	}
	if green == true {
		main_color = "green"
	}
	if orange == true {
		main_color = "orange"
	}
	if grey == true {
		main_color = "grey"
	}
	if silver == true {
		main_color = "silver"
	}
	if brown == true {
		main_color = "brown"
	}
	if beige == true {
		main_color = "beige"
	}
	if purple == true {
		main_color = "purple"
	}
	if pink == true {
		main_color = "pink"
	}
	if gold == true {
		main_color = "gold"
	}

	main_color_cl := strings.TrimSpace(main_color) //remove last space

	return main_color_cl
}
func SetSizeFromTitle(title string) string {
	ItemSize := ""
	title_lc := strings.ToLower(title)

	small := strings.Contains(title_lc, "small")
	medium := strings.Contains(title_lc, "medium")
	large := strings.Contains(title_lc, "large")
	extra_large := strings.Contains(title_lc, "extra large")

	small = strings.Contains(title_lc, "sml")
	medium = strings.Contains(title_lc, "med")
	large = strings.Contains(title_lc, "lg")
	extra_large = strings.Contains(title_lc, "xl")

	if small == true {
		ItemSize = "small"
	}
	if medium == true {
		ItemSize = "medium"
	}
	if large == true {
		ItemSize = "large"
	}
	if extra_large == true {
		ItemSize = "extra-large"
	}
	color.Cyan.Println(ItemSize)

	return ItemSize
}
func EncodeToString(EncStr string) string {
	data := []byte(EncStr)
	str := base64.StdEncoding.EncodeToString(data)
	return str
	// Output:
	// YW55ICsgb2xkICYgZGF0YQ==
}
