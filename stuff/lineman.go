package stuff

import "tcell"

func LineEnter(scr Bruh) Bruh {
	//convert cursor of screen to cursor of string

	// make temp array of strings
	temp := make([]string, len(scr.Lines)+1)
	for i := 0; i < len(temp); i++ {
		if i < scr.YCursor {
			temp[i] = scr.Lines[i]
		} else if i == scr.YCursor {
			temp[i] = scr.Lines[i][:scr.XCursor]
		} else if i == (scr.YCursor + 1) {
			temp[i] = scr.Lines[i-1][scr.XCursor:]
		} else {
			temp[i] = scr.Lines[i-1]
		}
	}
	scr.YCursor++
	scr.XCursor = 0
	scr.Lines = temp
	return scr
}
func Backspace(scr Bruh) Bruh {
	if scr.XCursor > 0 {

		temp1 := scr.Lines[scr.YCursor][:scr.XCursor-1] //everything before cursor

		temp2 := scr.Lines[scr.YCursor][scr.XCursor:] //everything after cursor
		scr.Lines[scr.YCursor] = temp1 + temp2        //combine them
		scr.XCursor--                                 //move cursor back
	} else if scr.XCursor == 0 && scr.YCursor > 0 { //if we are at the beginning of a line
		temp := len(scr.Lines[scr.YCursor])                                          //length of the line we are on
		scr.Lines[scr.YCursor-1] = scr.Lines[scr.YCursor-1] + scr.Lines[scr.YCursor] //combine the lines
		for i := scr.YCursor; i < len(scr.Lines); i++ {                              //move all lines up
			if i == len(scr.Lines)-1 { //if we are at the last line
				scr.Lines[i] = "" //set it to empty
			} else { //if we are not at the last line
				scr.Lines[i] = scr.Lines[i+1] //move the line up
			}
		}
		scr.XCursor = len(scr.Lines[scr.YCursor-1]) - temp //move cursor to end of previous line
		scr.YCursor--                                      //move cursor up

	}
	return scr
}
func Insert(scr Bruh, ev *tcell.EventKey) Bruh {
	scr.Lines[scr.YCursor] = scr.Lines[scr.YCursor][:scr.XCursor] + string(ev.Rune()) + scr.Lines[scr.YCursor][scr.XCursor:]
	scr.XCursor++
	return scr
}
func Delete(scr Bruh) Bruh {
	if scr.XCursor < len(scr.Lines[scr.YCursor]) {
		//we are not erasing what the cursor is on, but what is after it
		temp1 := scr.Lines[scr.YCursor][:scr.XCursor]
		temp2 := scr.Lines[scr.YCursor][scr.XCursor:]
		scr.Lines[scr.YCursor] = temp1 + temp2
	}
	return scr
}
func DeleteLine(scr Bruh) Bruh {
	for i := scr.YCursor; i < len(scr.Lines); i++ { //move all lines up
		if i == len(scr.Lines)-1 { //if we are at the last line
			scr.Lines[i] = "" //set it to empty
		} else { //if we are not at the last line
			scr.Lines[i] = scr.Lines[i+1] //move the line up
		}
	}
	return scr
}

// split a line into words
func SplitLine(scr Bruh) []string {
	var words []string
	var word string
	for _, char := range scr.Lines[scr.YCursor] {
		if char == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	words = append(words, word)
	return words
}
