//Explanation:
//Notifier Interface: Defines the contract for sending notifications. This decouples the NotificationService from specific notification methods.
//Concrete Implementations (EmailNotifier, SMSNotifier): Implement the Notifier interface, allowing interchangeable notification strategies.
//NotificationService Class: Uses the Notifier interface to send notifications. The specific implementation can be swapped easily, promoting low coupling.
//Client Code: Demonstrates how to switch between different notifiers by simply changing the implementation without modifying the NotificationService.

package main

import "fmt"

// Notifier interface defines the contract for sending notifications
type Notifier interface {
	SendNotification(message string)
}

// EmailNotifier is a concrete implementation of the Notifier interface
type EmailNotifier struct{}

// SendNotification sends an email notification
func (e *EmailNotifier) SendNotification(message string) {
	fmt.Printf("Sending email with message: %s\n", message)
}

// SMSNotifier is another concrete implementation of the Notifier interface
type SMSNotifier struct{}

// SendNotification sends an SMS notification
func (s *SMSNotifier) SendNotification(message string) {
	fmt.Printf("Sending SMS with message: %s\n", message)
}

// NotificationService uses the Notifier interface to send notifications
type NotificationService struct {
	notifier Notifier
}

// NewNotificationService creates a NotificationService with the specified Notifier
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

// Notify uses the Notifier to send a message
func (n *NotificationService) Notify(message string) {
	n.notifier.SendNotification(message)
}

func main() {
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{}

	// Create a notification service with EmailNotifier
	service := NewNotificationService(emailNotifier)
	service.Notify("Hello via Email")

	// Switch to SMSNotifier
	service = NewNotificationService(smsNotifier)
	service.Notify("Hello via SMS")
}

//OUTPUT
//
//Sending email with message: Hello via Email
//Sending SMS with message: Hello via SMS
