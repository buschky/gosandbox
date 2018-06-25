package kata

import s "strings"

func RepeatStr(repetitions int, value string) string {
/*
       var s string = ""
      for i := 0; i<repetitions; i++ {

        s = value[i] 
      }
*/
    return s.Repeat(value, repetitions)
}
