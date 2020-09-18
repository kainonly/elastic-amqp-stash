package session

import "errors"

func (c *Session) Nack(queue string, receipt string) (err error) {
	msg := c.receipt.Get(receipt)
	if msg == nil {
		return errors.New("the receipt has expired")
	}
	if msg.Queue != queue {
		return errors.New("the receipt verification is incorrect")
	}
	err = msg.Delivery.Nack(false, false)
	if err != nil {
		return
	}
	return
}