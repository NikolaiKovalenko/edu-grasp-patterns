//Explanation:
//Article Struct: Handles the data and functionality related to an article's content. It is responsible only for formatting its content, achieving high cohesion.
//Comment Struct: Manages individual comments with a focus on displaying content. It encapsulates the logic related to comment functionalities.
//CommentService Struct: Manages a collection of comments. It focuses on operations related to handling multiple comments, such as adding and displaying them.
//Client Code: Demonstrates the usage of these classes, highlighting their focused responsibilities.

package main

import (
	"fmt"
)

// Article is responsible for handling blog article content
type Article struct {
	Title, Body string
}

// Format formats the article content
func (a *Article) Format() string {
	return fmt.Sprintf("Title: %s\n\n%s", a.Title, a.Body)
}

// Comment is responsible for handling a single comment
type Comment struct {
	Author, Content string
}

// Display formats the comment for display
func (c *Comment) Display() string {
	return fmt.Sprintf("%s: %s", c.Author, c.Content)
}

// CommentService is responsible for managing multiple comments
type CommentService struct {
	comments []Comment
}

// AddComment adds a comment to the service
func (cs *CommentService) AddComment(comment Comment) {
	cs.comments = append(cs.comments, comment)
}

// DisplayComments displays all comments
func (cs *CommentService) DisplayComments() {
	for _, comment := range cs.comments {
		fmt.Println(comment.Display())
	}
}

func main() {
	// Create a new article
	article := &Article{
		Title: "High Cohesion in Software Design",
		Body:  "High cohesion refers to how closely related and focused the responsibilities of a class or function are.",
	}

	// Format and display the article
	fmt.Println(article.Format())

	// Initialize the comment service
	commentService := &CommentService{}

	// Add comments to the article
	commentService.AddComment(Comment{Author: "Alice", Content: "Great article!"})
	commentService.AddComment(Comment{Author: "Bob", Content: "Very informative."})

	// Display comments
	commentService.DisplayComments()
}

//OUTPUT
//
//Title: High Cohesion in Software Design
//
//High cohesion refers to how closely related and focused the responsibilities of a class or function are.
//Alice: Great article!
//Bob: Very informative.
