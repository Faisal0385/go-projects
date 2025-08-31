package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
)

// serveFile handles static files and custom 404
func serveFile(w http.ResponseWriter, r *http.Request) {
	baseDir := "my-website"
	path := filepath.Join(baseDir, r.URL.Path)

	// Default index.html
	if r.URL.Path == "/" {
		http.ServeFile(w, r, filepath.Join(baseDir, "index.html"))
		return
	}

	// Check if requested file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(baseDir, "404.html"))
		return
	}

	// Serve the file
	http.ServeFile(w, r, path)
}

// handleContact handles POST requests from the contact form
func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/contact.html", http.StatusSeeOther)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Send email
	sendEmail(name, email, message)

	// Redirect to thank-you page
	http.Redirect(w, r, "/thankyou.html", http.StatusSeeOther)
}

// sendEmail uses SMTP to send form submissions
func sendEmail(name, email, message string) {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com" // Gmail SMTP
	smtpPort := "587"
	sender := "faisal03gmit@gmail.com"
	password := "adki nopk ywqj thpr"
	receiver := "faisal03gmit@gmail.com"

	// Subject
	subject := "New Contact Form Submission"

	// HTML body
	body := fmt.Sprintf(`
		<html>
			<head>
			<style>
				body { font-family: Arial, sans-serif; line-height: 1.6; }
				h2 { color: #2a3693; }
				p { margin: 5px 0; }
				.label { font-weight: bold; }
			</style>
			</head>
			<body>
			<h2>New Contact Form Message</h2>
			<p><span class="label">Name:</span> %s</p>
			<p><span class="label">Email:</span> %s</p>
			<p><span class="label">Message:</span><br>%s</p>
			</body>
		</html>
		`, name, email, message)

	// Combine headers and body
	msg := []byte("To: " + receiver + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// SMTP authentication
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{receiver}, msg)
	if err != nil {
		log.Println("❌ Failed to send email:", err)
	} else {
		log.Println("✅ Contact form email sent successfully")
	}
}

func main() {
	// Serve static files
	http.HandleFunc("/", serveFile)

	// Handle contact form submission
	http.HandleFunc("/contact", handleContact)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
