package service

import (
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-transfer/config"
	"github.com/sirupsen/logrus"
)

func SubscribeEurRaw() {
	logrus.Info("subscribing to eur topic")
	_ = kfklib.Subscribe(config.EurBuildConfigInstance.Kafka.Group, EurBuildHandle, []string{config.EurBuildConfigInstance.Kafka.Topic})
}

func SubscribeGiteeIssue() {
	logrus.Info("subscribing to issue topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.Issue.Group, GiteeIssueHandle, []string{config.GiteeConfigInstance.Issue.Topic})
}

func SubscribeGiteePush() {
	logrus.Info("subscribing to push topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.Push.Group, GiteePushHandle, []string{config.GiteeConfigInstance.Push.Topic})
}

func SubscribeGiteePr() {
	logrus.Info("subscribing to pr topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.PR.Group, GiteePrHandle, []string{config.GiteeConfigInstance.PR.Topic})
}

func SubscribeGiteeNote() {
	logrus.Info("subscribing to note topic")
	_ = kfklib.Subscribe(config.GiteeConfigInstance.Note.Group, GiteeNoteHandle, []string{config.GiteeConfigInstance.Note.Topic})
}

func SubscribeOpenEulerMeeting() {
	logrus.Info("subscribing to openEuler meeting topic")
	_ = kfklib.Subscribe(config.MeetingConfigInstance.Kafka.Group, OpenEulerMeetingHandle, []string{config.MeetingConfigInstance.Kafka.Topic})
}

func SubscribeCVERaw() {
	logrus.Info("subscribing to cve topic")
	_ = kfklib.Subscribe(config.CveConfigInstance.Kafka.Group, CVEHandle, []string{config.CveConfigInstance.Kafka.Topic})
}
