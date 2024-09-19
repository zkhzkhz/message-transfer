package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/message-transfer/common/kafka"
	"github.com/opensourceways/message-transfer/config"
	"github.com/opensourceways/message-transfer/models/dto"
	"github.com/opensourceways/message-transfer/utils"
)

func handle(raw dto.Raw, cfg kafka.ConsumeConfig) error {
	time.Sleep(utils.GetConsumeSleepTime())
	event := raw.ToCloudEventByConfig(cfg.Topic)
	if event.ID() == "" {
		return nil
	}
	err := event.SaveDb()
	if err != nil {
		logrus.Errorf("saveDb failed, err:%v", err)
		return err
	}
	kafkaSendErr := kafka.SendMsg(cfg.Publish, &event)
	if kafkaSendErr != nil {
		return kafkaSendErr
	}
	return nil
}

func CVEHandle(payload []byte, _ map[string]string) error {
	var raw dto.CVEIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	rawMap := raw.ToMap()
	return handle(rawMap, config.CveConfigInstance.Kafka)
}

func GiteeIssueHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeIssueRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	raw.SigGroupName = sigGroupName
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Issue)
}

func GiteePushHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePushRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	raw.SigGroupName = sigGroupName
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Push)
}

func GiteePrHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteePrRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	raw.SigGroupName = sigGroupName
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.PR)
}

func GiteeNoteHandle(payload []byte, _ map[string]string) error {
	var raw dto.GiteeNoteRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	sigGroupName, err := utils.GetRepoSigInfo(raw.Repository.Name)
	if err != nil {
		return err
	}
	raw.SigGroupName = sigGroupName
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.GiteeConfigInstance.Note)
}

func EurBuildHandle(payload []byte, _ map[string]string) error {
	var raw dto.EurBuildMessageRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	raw.SourceGroup = fmt.Sprintf("%s/%s", raw.Body.Owner, raw.Body.Copr)
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.EurBuildConfigInstance.Kafka)
}

func OpenEulerMeetingHandle(payload []byte, _ map[string]string) error {
	var raw dto.OpenEulerMeetingRaw
	msgBodyErr := json.Unmarshal(payload, &raw)
	if msgBodyErr != nil {
		return msgBodyErr
	}
	raw.MeetingStartTime = raw.Msg.Date + raw.Msg.Start
	raw.MeetingEndTime = raw.Msg.Date + raw.Msg.End
	rawMap := dto.StructToMap(raw)
	return handle(rawMap, config.MeetingConfigInstance.Kafka)
}
