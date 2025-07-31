package kafkaapi

import (
	"context"
	"fmt"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
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
			msg, err := api.consumer.ReadMessage(-1)
			if err != nil {
				api.logger.Send("error", supportingfunctions.CustomError(err).Error())
			}

			topic := msg.TopicPartition.Topic

			subjectType := "undefined_type"
			topicValue, ok := api.topics[*topic]
			if ok {
				subjectType = topicValue
			}

			api.chFromModule <- SettingsChanOutput{
				SubjectType: subjectType,
				Data:        msg.Value,
			}
		}
	}
}
