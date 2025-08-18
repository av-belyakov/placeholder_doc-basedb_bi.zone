package kafkaapi

import (
	"context"
	"fmt"
)

// topicsHandler обработчик топиков (подписок)
func (api *kafkaApiModule) topicsHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("func 'topicsHandler', закрываем канал на отправку из модуля")

			//закрываем канал на отправку из модуля
			close(api.chFromModule)

			return

		default:
			fmt.Println("Default msg")
			/*
				msg, err := api.consumer.ReadMessage(-1)
				if err != nil {
					api.logger.Send("error", supportingfunctions.CustomError(err).Error())
				}

				topic := msg.TopicPartition.Topic

				subjectType := "undefined_type"
				topicKey, ok := supportingfunctions.SearchValue(api.topics, *topic)
				if ok {
					subjectType = topicKey
				}

				api.chFromModule <- SettingsChanOutput{
					SubjectType: subjectType,
					Data:        msg.Value,
				}
			*/
		}
	}
}
