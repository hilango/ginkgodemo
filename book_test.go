package ginkgodemo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//. "github.com/hilango/ginkgodemo"
)

type Book struct{
	Title string
	Author string
	Pages	int
}

func (b *Book) CategoryByLength() (string) {
	if(b.Pages<300) {
		return "SHORT STORY"
	}else {
		return "NOVEL"
	}
}

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("根据页数分类", func() {
		Context("大于等300页", func() {
			It("应该返回小说", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("小于300页", func() {
			It("应该返回短故事", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})

