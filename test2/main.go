package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	CommentId      int        `json:"commentId"`
	CommentContent string     `json:"commentContent"`
	Replies        []*Comment `json:"replies"`
}

func main() {
	// Data JSON yang diberikan
	jsonData := `[ 
		{
			"commentId": 1,
			"commentContent": "Hai",
			"replies": [
				{
					"commentId": 11,
					"commentContent": "Hai juga",
					"replies": [
						{
							"commentId": 111,
							"commentContent": "Haai juga hai jugaa"
						},
						{
							"commentId": 112,
							"commentContent": "Haai juga hai jugaa"
						}
					]
				},
				{
					"commentId": 12,
					"commentContent": "Hai juga",
					"replies": [
						{
							"commentId": 121,
							"commentContent": "Haai juga hai jugaa"
						}
					]
				}
			]
		},
		{
			"commentId": 2,
			"commentContent": "Halooo"
		}
	]`

	// Parsing JSON ke dalam slice dari struct Comment
	var comments []*Comment
	err := json.Unmarshal([]byte(jsonData), &comments)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Hitung total komentar termasuk balasannya
	totalComments := countComments(comments)
	fmt.Println("Total komentar:", totalComments)
}

// Fungsi rekursif untuk menghitung total komentar termasuk balasannya
func countComments(comments []*Comment) int {
	count := len(comments)

	for _, comment := range comments {
		count += countComments(comment.Replies)
	}

	return count
}
