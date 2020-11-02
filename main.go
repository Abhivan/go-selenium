// package selenium_test

package main

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/tebeka/selenium"
)

func main() {

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	// wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:4444/wd/hub"))
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://selenium-hub:4444/wd/hub"))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to URL.
	if err := wd.Get("https://www.vfsvisaservice.com/IHC-UK-Appt/AppScheduling/AppWelcome.aspx?P=cCiy6xeqlBWf0MSvlUERSCDhdFkas/mFaceUVLcp3A4="); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	schAppElem, err := wd.FindElement(selenium.ByCSSSelector, "#ctl00_plhMain_lnkSchApp")
	if err != nil {
		panic(err)
	}
	// Click Schedule appointment.
	if err := schAppElem.Click(); err != nil {
		panic(err)
	}

	vacElem, err := wd.FindElement(selenium.ByCSSSelector, "#ctl00_plhMain_cboVAC > option:nth-child(3)")
	if err != nil {
		panic(err)
	}

	if err := vacElem.Click(); err != nil {
		panic(err)
	}

	serElem, err := wd.FindElement(selenium.ByCSSSelector, "#ctl00_plhMain_cboVisaCategory > option:nth-child(2)")
	if err != nil {
		panic(err)
	}

	// Select London-Hounslow.
	if err := serElem.Click(); err != nil {
		panic(err)
	}

	msgDiv, err := wd.FindElement(selenium.ByCSSSelector, "#ctl00_plhMain_lblMsg")
	if err != nil {
		panic(err)
	}

	msg, err := msgDiv.Text()
	if err != nil {
		panic(err)
	}

	fmt.Printf(msg)

	// Send email alert if the slots are availble
	if strings.Contains(msg, "Sorry") == true {
		fmt.Printf("Not sending alerts")
	} else {
		err := SendMail("127.0.0.1:25", (&mail.Address{"from name", "vfs-alertt@localhost"}).String(), "Alert- Appointments Now available", "Now appointments are available", []string{(&mail.Address{"to name", "yskci18971@adeata.com"}).String()})
		if err != nil {
			panic(err)
		}
	}

}

//ex: SendMail("127.0.0.1:25", (&mail.Address{"from name", "from@example.com"}).String(), "Email Subject", "message body", []string{(&mail.Address{"to name", "to@example.com"}).String()})
func SendMail(addr, from, subject, body string, to []string) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
