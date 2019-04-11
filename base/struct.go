package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books       
   var Book2 Books        

   /* book 1 desc */
   Book1.title = "Go lang"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go lang"
   Book1.book_id = 6495407

   /* book 2 desc */
   Book2.title = "Python"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python"
   Book2.book_id = 6495700

   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)


   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
