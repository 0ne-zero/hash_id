# Hash Identifier
## General Information
- You can pass a hash to this module and it gives back you possible algorithms of the hash

## Usage
- First get the module
```bash
$ go get github.com/0ne-zero/hash_id
```
- Then write a program that uses the module
```golang
import "github.com/0ne-zero/hash_id"
import "fmt"
func main(){
  hash := "73cb3858a687a8494ca3323053016282f3dad39d42cf62ca4e79dda2aac7d9ac"
  
  // Here you go
  algos := hash_id.Identify(hash)
  
  // Output: [SHA256 HAVAL256]
  fmt.Println(algos)
}
```

## Contributions
- It's one of those projects that needs your contributions.

## Contact
- Created by [Pouria Khakpour](https://github.com/0ne-zero) - feel free to contact me!
- You can reach me by pouria.khakpour9909@gmail.com
