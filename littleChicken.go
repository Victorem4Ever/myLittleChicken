package main

import (
	"fmt"
   "os/exec"
   "os"
   "strings"
   "strconv"
)

func main() {
	chicken := `    \\
    ( °>
 \\_//)
  \_/_)
   _|_`
   game_datas := [...]string{"", "", "", "", "", "", "", "", ""}
   game(game_datas, chicken)
}

func game(datas [9]string, chicken string){
   running := true
   var input string
   help_command := `Commands :
   help : Show this command
   exit : Exit the game
   eat : Give food to the chicken
   rename : Rename the chicken`

   name := "Chicky"

   stats := [...]int{15, 0, 0, 0, 5}

	for running {
      clear()

      fmt.Println(chicken)
      fmt.Println(format_stats(stats, name))

      fmt.Printf("Command :\n> ")
      fmt.Scanln(&input)

      input = strings.ToLower(input)

      if input == "exit"{
         running = false
         fmt.Printf("\nGood bye!\nPress enter to exit ...")
         fmt.Scanln()
         continue
   
      } else if input == "help"{

         fmt.Println(help_command)
         fmt.Printf("Press enter to continue ...")
         fmt.Scanln()
         continue

      } else if input == "eat"{
         if stats[0] >= 20{
            fmt.Printf("%v is going to explose !!!", name)
            stats[0] += 1
            
            if stats[0] == 25{
               fmt.Printf("You killed %v", name)
               stats[4] = 0
            } else{
               stats[4] -= 1
            }
         } else if stats[0] == 19{
            fmt.Printf("%v is now totally full!", name)
            stats[0] += 1
            if stats[4] < 10{
               stats[4] += 1
            } 

         } else{
               stats[0] += 1
         }

         fmt.Printf("\nPress enter to continue ...")
         fmt.Scanln()
         continue
      } else if input == "rename" {
         fmt.Printf("\nEnter the new name :\n> ")
         fmt.Scanln(&name)
         fmt.Printf("\nPress enter to continue ...")
         fmt.Scanln()
         continue
      } else{
         fmt.Printf("\nUnknow command\nPress enter to continue ...")
         fmt.Scanln()
      }
   }
}

func clear(){

   // Windows = cls - Linux = clear
   cmd := exec.Command("cmd", "/c", "cls")
   cmd.Stdout = os.Stdout
   cmd.Run()
}

func format_stats(stats [5]int, name string) string{
   return "Food : " + strconv.Itoa(stats[0]) + "\nAlive while : " + strconv.Itoa(stats[1]) + " mins.\nDead since : " + strconv.Itoa(stats[2]) + " mins.\nNazi : " + strconv.Itoa(stats[3]) + "%\nHearts : " + strconv.Itoa(stats[4]) + "\nName : " + name
}