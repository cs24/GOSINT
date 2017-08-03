package main

import (
	"fmt"
)

//PrintLogo is a simple function for printing the framework name called upon start.
func PrintLogo() {
	fmt.Print(`
          _____                  _______                  _____                   _____                   _____            _____          
         /\    \                /::\    \                /\    \                 /\    \                 /\    \          /\    \         
        /::\    \              /::::\    \              /::\    \               /::\    \               /::\____\        /::\    \        
       /::::\    \            /::::::\    \            /::::\    \              \:::\    \             /::::|   |        \:::\    \       
      /::::::\    \          /::::::::\    \          /::::::\    \              \:::\    \           /:::::|   |         \:::\    \      
     /:::/\:::\    \        /:::/~~\:::\    \        /:::/\:::\    \              \:::\    \         /::::::|   |          \:::\    \     
    /:::/  \:::\    \      /:::/    \:::\    \      /:::/__\:::\    \              \:::\    \       /:::/|::|   |           \:::\    \    
   /:::/    \:::\    \    /:::/    / \:::\    \     \:::\   \:::\    \             /::::\    \     /:::/ |::|   |           /::::\    \   
  /:::/    / \:::\    \  /:::/____/   \:::\____\  ___\:::\   \:::\    \   ____    /::::::\    \   /:::/  |::|   | _____    /::::::\    \  
 /:::/    /   \:::\ ___\|:::|    |     |:::|    |/\   \:::\   \:::\    \ /\   \  /:::/\:::\    \ /:::/   |::|   |/\    \  /:::/\:::\    \ 
/:::/____/  ___\:::|    |:::|____|     |:::|    /::\   \:::\   \:::\____/::\   \/:::/  \:::\____/:: /    |::|   /::\____\/:::/  \:::\____\
\:::\    \ /\  /:::|____|\:::\    \   /:::/    /\:::\   \:::\   \::/    \:::\  /:::/    \::/    \::/    /|::|  /:::/    /:::/    \::/    /
 \:::\    /::\ \::/    /  \:::\    \ /:::/    /  \:::\   \:::\   \/____/ \:::\/:::/    / \/____/ \/____/ |::| /:::/    /:::/    / \/____/ 
  \:::\   \:::\ \/____/    \:::\    /:::/    /    \:::\   \:::\    \      \::::::/    /                  |::|/:::/    /:::/    /          
   \:::\   \:::\____\       \:::\__/:::/    /      \:::\   \:::\____\      \::::/____/                   |::::::/    /:::/    /           
    \:::\  /:::/    /        \::::::::/    /        \:::\  /:::/    /       \:::\    \                   |:::::/    /\::/    /            
     \:::\/:::/    /          \::::::/    /          \:::\/:::/    /         \:::\    \                  |::::/    /  \/____/             
      \::::::/    /            \::::/    /            \::::::/    /           \:::\    \                 /:::/    /                       
       \::::/    /              \::/____/              \::::/    /             \:::\____\               /:::/    /                        
        \::/____/                ~~                     \::/    /               \::/    /               \::/    /                         
                                                         \/____/                 \/____/                 \/____/                          
                                                                                                                                          
`)
}