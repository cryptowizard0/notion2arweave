package utils

import (
	arTypes "github.com/everFinance/goar/types"
	"github.com/spf13/viper"
)

func MakeTags(typ, sign, content string) []arTypes.Tag {
	var tags []arTypes.Tag
	tags = append(tags, []arTypes.Tag{
		// Base tags
		{Name: "Content-Type", Value: "application/json"},
		{Name: "App-Name", Value: viper.GetString("appname")},
		{Name: "App-Version", Value: viper.GetString("version")},
		// App tags
		{Name: "type", Value: typ},
		{Name: "sign", Value: sign},
		{Name: "content", Value: content},
	}...)

	return tags
}

func MakeImageTags(sign string) []arTypes.Tag {
	var tags []arTypes.Tag
	tags = append(tags, []arTypes.Tag{
		// Base tags
		{Name: "Content-Type", Value: "image/jpeg"},
		{Name: "App-Name", Value: viper.GetString("appname")},
		{Name: "App-Version", Value: viper.GetString("version")},
		// App tags
		{Name: "type", Value: "image"},
		{Name: "sign", Value: sign},
	}...)

	return tags
}

func GetTagValue(name string, tags []arTypes.Tag) string {
	for _, tag := range tags {
		if tag.Name == name {
			return tag.Value
		}
	}
	return ""
}
