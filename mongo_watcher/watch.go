package mongo_watcher

import (
	"context"
	"fmt"
	"log"
	"luckperms-notifier/config"
	"luckperms-notifier/endpoints"
	"luckperms-notifier/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChangeEvent struct {
    OperationType string `bson:"operationType"`
    DocumentKey struct {
        ID interface{} `bson:"_id"`
    } `bson:"documentKey"`
    FullDocument struct {
        Description string `bson:"description"`
        Source struct {
            Name string `bson:"name"`
        } `bson:"source"`
        Target struct {
            Type string `bson:"type"`
            Name string `bson:"name"`
        } `bson:"target"`
    } `bson:"fullDocument"`
} 

var (
    TITLE string = "New action"
    COLOR string = "2353957"
    INLINE bool = true
    FOOTER string = "github.com/IdanKoblik/Luckperms-notifier"
) 

func WatchCollection(collection *mongo.Collection, options *options.ChangeStreamOptions) {
    changeStream, err := collection.Watch(context.Background(), mongo.Pipeline{}, options); if err != nil {
        log.Fatalf("Error opening change stream: %v\n", err)
        return
    }

    defer changeStream.Close(context.Background())

    for changeStream.Next(context.Background()) {
        var changeEvent ChangeEvent

        if err := changeStream.Decode(&changeEvent); err != nil {
            log.Println("Error decoding change event:", err)
            continue
        }

        if changeEvent.OperationType == "insert" {
            webhookURL, err := config.GetURL(); if err != nil {
                log.Fatalf("Error getting the url of the webhook: %v\n", err)
                return
            }

            sendEmbed(&changeEvent, webhookURL)
        }
    }

    if err := changeStream.Err(); err != nil {
        log.Println("Error from change stream:", err)
        return
    }
}

func getEmbed(changeEvent *ChangeEvent, webhookURL string) (utils.Embed, error) {
    thumbnail, err := getThumbnail(webhookURL); if err != nil {
        return utils.Embed{}, err
    }

    description := "**" + changeEvent.FullDocument.Description + "**"
    fields := []utils.Field{}
    fields = append(fields, getTypeField(changeEvent))
    fields = append(fields, getNameField(changeEvent))
    fields = append(fields, getSourceField(changeEvent))

    return utils.Embed {
        Title: &TITLE,
        Description: &description,
        Fields: &fields,
        Color: &COLOR,
        Thumbnail: &thumbnail,
        Footer: &utils.Footer{
            Text: &FOOTER,
        },
    }, nil
}

func sendEmbed(changeEvent *ChangeEvent, webhookURL string) {
    hook, err := endpoints.Fetch(webhookURL); if err != nil {
        log.Fatalf("Error getting the username of the webhook: %v\n", err)
        return
    }

    embed, err := getEmbed(changeEvent, webhookURL); if err != nil {
        log.Fatalf("%v\n", err)
        return
    }

    message := utils.Message{
        Username: &hook.Name,
        Embeds:   &[]utils.Embed{embed},
    }

    message.SendMessage(webhookURL)
}

func getNameField(changeEvent *ChangeEvent) (utils.Field) {
    typeFieldName := "Type"
    return utils.Field {
        Name: &typeFieldName,
        Value: &changeEvent.FullDocument.Target.Type,
        Inline: &INLINE,
    }
}

func getTypeField(changeEvent *ChangeEvent) (utils.Field) {
    nameFieldName := "Name"
    return utils.Field {
        Name: &nameFieldName,
        Value: &changeEvent.FullDocument.Target.Name,
        Inline: &INLINE,
    }
}

func getThumbnail(webhookURL string) (utils.Thumbnail, error) {
    hook, err := endpoints.Fetch(webhookURL); if err != nil {
        return utils.Thumbnail{}, err
    }

    url := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", hook.Id, hook.Avatar)
    return utils.Thumbnail{Url: &url}, nil
}

func getSourceField(changeEvent *ChangeEvent) (utils.Field) {
    source := changeEvent.FullDocument.Source.Name
    sourceFieldName := "Source"
    inline := true

    return utils.Field{
        Name: &sourceFieldName,
        Value: &source,
        Inline: &inline,
    }
}