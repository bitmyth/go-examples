// https://github.com/Sirupsen/logrus
package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears A walrus appears A walrus appears")

	event:="test event"
	topic:="test topic"
	key:="test key"

	log.WithFields(log.Fields{
		"event": event,
		"topic": topic,
		"key": key,
	}).Fatal("Failed to send event")
}
