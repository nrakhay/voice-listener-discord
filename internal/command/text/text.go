package text

import (
	"bytes"
	"fmt"
	"log/slog"

	"github.com/nrakhay/ONEsports/internal/discord"
	"github.com/nrakhay/ONEsports/internal/service/s3"

	"github.com/bwmarrin/discordgo"
)

func SendVoiceRecordingToTextChannel(channelId string, channel string, key string) {
	fileBuffer, err := s3.RetrieveFileFromS3(key)

	if err != nil {
		slog.Error("Failed to retrieve file from S3: ", err)
		return
	}

	file := &discordgo.File{
		Name:   key,
		Reader: bytes.NewReader(fileBuffer),
	}
	messageSend := discordgo.MessageSend{
		Content: fmt.Sprintf("Recording from voice channel: **%s**", channel),
		Files:   []*discordgo.File{file},
	}

	_, err = discord.Session.ChannelMessageSendComplex(channelId, &messageSend)
	if err != nil {
		slog.Error("Failed to send file to text channel: ", err)
		return
	}

	slog.Info("Successfully sent file to text channel.", "Filename", key)
}
