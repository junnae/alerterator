package receivers

import (
	"testing"

	"alerterator/utils"

	"alerterator/controllers/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestReceivers(t *testing.T) {
	t.Run("Validating that receivers are created correctly", func(t *testing.T) {
		receiver := createReceiver(fixtures.AlertResource)
		assert.Equal(t, utils.GetCombinedName(fixtures.AlertResource), receiver.Name)
		assert.Len(t, receiver.EmailConfigs, 1)
		assert.Len(t, receiver.SlackConfigs, 1)

		receivers := fixtures.AlertResource.Spec.Receivers
		assert.Equal(t, receivers.Email.To, receiver.EmailConfigs[0].To)
		assert.Equal(t, receivers.Email.SendResolved, receiver.EmailConfigs[0].SendResolved)

		assert.Equal(t, receivers.Slack.Channel, receiver.SlackConfigs[0].Channel)
		assert.Equal(t, receivers.Slack.PrependText, fixtures.AlertResource.Spec.Receivers.Slack.PrependText)

		assert.Equal(t, receivers.SMS.SendResolved, receiver.WebhookConfigs[0].SendResolved)
	})

	t.Run("Valider at send_resolved for e-post blir beholdt", func(t *testing.T) {
		alert := fixtures.AlertResource
		alert.Spec.Receivers.Email.SendResolved = true
		receiver := createReceiver(alert)
		assert.True(t, receiver.EmailConfigs[0].SendResolved)
	})

	t.Run("Valider at send_resolved for sms blir beholdt", func(t *testing.T) {
		alert := fixtures.AlertResource
		alert.Spec.Receivers.SMS.SendResolved = false
		receiver := createReceiver(alert)
		assert.False(t, receiver.WebhookConfigs[0].SendResolved)
	})

	t.Run("Valider at user key blir satt", func(t *testing.T) {
		alert := fixtures.AlertResource
		receiver := createReceiver(alert)
		assert.Equal(t, receiver.PushoverConfigs[0].UserKey, alert.Spec.Receivers.Pushover.UserKey)
	})
}
