package main

import "fmt"

// Definisikan struktur IComment
type IComment struct {
	CommentId      int
	CommentContent string
	Replies        []IComment
}

// Fungsi untuk menghitung total komentar dan semua balasannya
func CountTotalComments(comments []IComment) int {
	totalComments := 0

	// Iterasi melalui setiap komentar
	for _, comment := range comments {
		// Tambahkan 1 untuk setiap komentar
		totalComments++

		// Jika komentar memiliki balasan, rekursif hitung total komentar di balasan tersebut
		totalComments += CountTotalComments(comment.Replies)
	}

	return totalComments
}

func main() {
	// Deklarasi data komentar
	comments := []IComment{
		{
			CommentId:      1,
			CommentContent: "Hai",
			Replies: []IComment{
				{
					CommentId:      11,
					CommentContent: "Hai juga",
					Replies: []IComment{
						{
							CommentId:      111,
							CommentContent: "Haai juga hai jugaa",
						},
						{
							CommentId:      112,
							CommentContent: "Haai juga hai jugaa",
						},
					},
				},
				{
					CommentId:      12,
					CommentContent: "Hai juga",
					Replies: []IComment{
						{
							CommentId:      121,
							CommentContent: "Haai juga hai jugaa",
						},
					},
				},
			},
		},
		{
			CommentId:      2,
			CommentContent: "Halooo",
		},
	}

	// Hitung total komentar dan cetak hasilnya
	totalComments := CountTotalComments(comments)
	fmt.Printf("Total komentar adalah %d komentar.\n", totalComments)
}
