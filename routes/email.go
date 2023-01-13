package routes

// ตัวอย่างการใช้ text/template https://betterprogramming.pub/how-to-use-templates-in-golang-46194c677c7d
import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

type MessageData struct {
	HospitalNameTh string
	Username       string
	Password       string
	Permission     string
}

// ส่ง text และ value แบบไหนก็ได้ให้สัมพันธ์กัน
func setValueAndText(s string, val interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, val)
	return b.String()
}

// บังคับส่งฟอร์มแบบ struct นี้เท่านั้น
func setValueFromStruct(val MessageData) string {
	s := `<html>
	<body>
	  เรียน ผู้ใช้งาน
	  <table>
		<tr>
		  <td>{{.HospitalNameTh}}</td>
		</tr>
		<tr>
		  <td>URL :</td>
		  <td>
			<a href="https://google.com">https://google.com</a>
		  </td>
		</tr>
		<tr>
		  <td>Username :</td>
		  <td>{{.Username}}</td>
		</tr>
		<tr>
		  <td>Password :</td>
		  <td>{{.Password}}</td>
		</tr>
		<tr>
		  <td>Permission :</td>
		  <td>{{.Permission}}</td>
		</tr>
	  </table>
	  หากมีข้อสงสัยสามารถติดต่อมาได้ที่ contact@xxx.co.th
	</body>
	  </html>`
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, val)
	return b.String()
}

func RouterEmail(app fiber.Router) {
	app.Post("/gmail", func(c *fiber.Ctx) error {
		// Sender data
		from := "a5720187@gmail.com"
		password := "xeichqnipiuswqtvox"

		// Receiver email address
		to := []string{
			"boss57991@gmail.com",
		}

		// smtp server configuration
		host := "smtp.gmail.com"
		port := "587"

		// Authentication
		auth := smtp.PlainAuth("", from, password, host)

		// Create message
		subject := "Subject: แจ้ง Username/Password\n"
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" // จะใส่เมื่อส่ง body เป็น html
		val := MessageData{Username: "boss57991", HospitalNameTh: "โรงพยาบาลxxx", Password: "123456", Permission: "user"}
		body := setValueFromStruct(val)

		msg := []byte(subject + mimeHeaders + body)

		// Sending email
		err := smtp.SendMail(host+":"+port, auth, from, to, msg)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Email sent successfully")
		return nil
	})

	app.Post("/inet", func(c *fiber.Ctx) error {
		// Sender data
		sender := "julladith.kl@one.th"

		// Receiver email address
		to := []string{
			"boss57991@hotmail.com",
		}

		// smtp server configuration
		host := "mailtx.inet.co.th"
		port := "25"

		// Create message
		from := "From: julladith.kl@one.th\n"
		subject := "Subject: แจ้ง Username/Password สำหรับเข้าระบบ\n"
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" // จะใส่เมื่อส่ง body เป็น html
		val := MessageData{Username: "boss579912", HospitalNameTh: "โรงพยาบาลxxx", Password: "123456", Permission: "user"}
		body := setValueFromStruct(val)

		msg := []byte(from + subject + mimeHeaders + body)

		smtpDial, err := smtp.Dial(host + ":" + port)
		if err != nil {
			log.Panic(err)
		}

		// Config Sender
		if err = smtpDial.Mail(sender); err != nil {
			log.Panic(err)
		}

		// Config Reciever
		if err = smtpDial.Rcpt(to[0]); err != nil {
			log.Panic(err)
		}

		// Write Data
		w, err := smtpDial.Data()
		if err != nil {
			log.Panic(err)
		}

		_, err = w.Write([]byte(msg))
		if err != nil {
			log.Panic(err)
		}

		err = w.Close()
		if err != nil {
			log.Panic(err)
		}

		smtpDial.Quit()

		fmt.Println("Email sent successfully")

		return nil
	})
}
