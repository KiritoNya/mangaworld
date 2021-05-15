# mangaworld

[![GoDoc](https://godoc.org/github.com/gorilla/mux?status.svg)](https://pkg.go.dev/github.com/KiritoNya/mangaworld)
[![CircleCI](https://circleci.com/gh/gorilla/mux.svg?style=svg)](https://circleci.com/gh/kiritoNya/mangaworld)

Package `kiritoNya/mangaworld` implements many API that allow web scraping from the site [Mangaworld](https://www.mangaworld.io/).

It implements:

* Get information of manga and chapters.
* Get page url of chapters.
* Get trending manga and new releases.
* Search manga by creating queries and querying the site.
* Download the chapters.
---

* [Install](#install)
* [Examples](#examples)
* [Info](#info)
* [Query](#query)
* [New Chapters](#new-chapters-released)
* [Download](#download)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/kiritoNya/mangaworld
```

## Examples

Let's start by creating a manga object and a chapter object.

```go
func main() {
    
	//Creating manga object
	manga, err := mangaworld.NewManga("https://www.mangaworld.io/manga/395/citrus")
	if err != nil {
		log.Fatalln(err)
	}
	
	//Creating chapter object
	chapter, err := mangaworld.NewChapter("https://www.mangaworld.io/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbac0/1")
	if err != nil {
		log.Fatalln(err)
	}
	
}
```

We have seen how to create `Manga` type and `Chapter` type objects. There are other types of objects in the module that can be created in the same way. (See documentation)

### Info


Now let's see how to fill objects by invoking methods and getting manga or chapter data.

```go

    //Create object
    manga, _ := mangaworld.NewManga("https://www.mangaworld.io/manga/395/citrus")
    
    //Call getTitle method to get the title of manga
    err = manga.GetTitle()
    if err != nil {
        t.Error("Error to get title")
    }
    
    fmt.Println("TITLE:", manga.Title)
    fmt.Println("OBJECT:", manga)

```

### Query


Queries are useful for searching for manga using different filters.
Now let's see how they are created and how to apply some filters.

```go
func main() {
    
    //Create query
    q := mangaworld.NewQuery()
    
    //Serch by name
    q.SetMangaName("yagate kimi ni naru")
	
}
```

You can also apply multiple filters at the same time.

```go
func main() {
    
    //Create query
    q := mangaworld.NewQuery()

    //Name filter
    q.SetMangaName("citrus")
    
    //Genres filter
    q.SetGenres([]mangaworld.Genre{
    	mangaworld.Yuri, 
    	mangaworld.Shoujo_ai,
    })
    
    //Manga types filter
    q.SetMangaTypes([]mangaworld.Type{
    	mangaworld.Manga_type 
    })
    
    //Status fileter
    q.SetStatus([]State{ 
    	mangaworld.Finish 
    })
    
    //Year filter
    q.SetYears([]string{
    	"2012" 
    })
	
}
```

To start the query search

```go
func main() {
    
    //Create query
    q := mangaworld.NewQuery()
    
    //Serch by name
    q.SetMangaName("yagate kimi ni naru")
    
    //Do the query
    mangas, err := q.Do
    if err != nil {
    	log.Fatalln(err)
    }
	
    for _, manga := range mangas {
        fmt.Println(manga.Url)	
    }
    
}
```

### New Chapters Released

Another of the features made available by this module is to be able to get the latest chapters released. Let's see how to do it.

```go
func main() {
	
	//Get last 2 new chapters
	chaptersNew, err := mangaworld.ChaptersNew(2)
	if err != nil {
	    log.Fatalln(err)
	}
	
	for _, chapterNew := range chaptersNew {
		
		err := chapterNew.GetChapter()
		if err != nil {
			log.Fatalln(err)
		}
		
		err = chapterNew.GetManga()
		if err != nil {
			log.Fatalln(err)
		}
	}
	
}
```

### Download

We come to the part that will certainly interest you the most, namely the download of the chapters.

```go
func main() {
	
    //Create chapter
    chapter, err := mangaworld.NewChapter("https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbaac/1")
    if err != nil {
        t.Error(err)
    }

    //Download chapter in /home/<user>/new
    err = c.Download("/home/<user>/new")
    if err != nil {
    	log.Fatalln(err)
    }
    
}
```

## License

MIT licensed. See the LICENSE file for details.
