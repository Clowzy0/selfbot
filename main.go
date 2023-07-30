package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ecnepsnai/discord"
)

// token=MTEzNTIwNDk4MTYyMDg3MTE5OA.G78dO6.U7JOxWM6Hlfv2RHMMn9CnxDUfvLvX5yzDuO5I8

var (
	Token         string
	downloads     []string
	ID_1          string
	Server_ID_1   string
	Server_Name_1 string
	Name_1        string
	Folders_1     string
)

type name struct {
	ID          string
	Server_ID   string
	Server_Name string
	Name        string
	Folders     string
}

var names = []name{
	{ID: "953333224849956994", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Amber(Genshin Impact)", Folders: "Genshin,Genshin/Amber"},
	{ID: "952911848267718739", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Barbara(Genshin Impact)", Folders: "Genshin,Genshin/Barbara"},
	{ID: "953333669236453457", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Eula(Genshin Impact)", Folders: "Genshin,Genshin/Eula"},
	{ID: "953334063928836146", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Fischl(Genshin Impact)", Folders: "Genshin,Genshin/Fischl"},
	{ID: "953333083577401404", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Lisa(Genshin Impact)", Folders: "Genshin,Genshin/Lisa"},
	{ID: "953333169971675157", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Mona(Genshin Impact)", Folders: "Genshin,Genshin/Mona"},
	{ID: "953333610352635996", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Jean(Genshin Impact)", Folders: "Genshin,Genshin/Jean"},
	{ID: "954895410734960670", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Rosaria(Genshin Impact)", Folders: "Genshin,Genshin/Rosaria"},
	{ID: "954881022896533525", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Noelle(Genshin Impact)", Folders: "Genshin,Genshin/Noelle"},
	{ID: "953335813893136465", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Lumine(Genshin Impact)", Folders: "Genshin,Genshin/Lumine"},
	{ID: "961022085306667099", Server_ID: "952611412285485066", Server_Name: "Genshin_Hentai_Server", Name: "Sucrose(Genshin Impact)", Folders: "Genshin,Genshin/Sucrose"},
	{ID: "1115970003439714365", Server_ID: "946803206091059232", Server_Name: "Cumsluts_18+", Name: "trendy_hentai_videos", Folders: "trendy_hentai_vid"},
	{ID: "1121673875164504074", Server_ID: "946803206091059232", Server_Name: "Cumsluts_18+", Name: "hentai_images", Folders: "hentai_img"},
	{ID: "1115970121031225447", Server_ID: "946803206091059232", Server_Name: "Cumsluts_18+", Name: "porn_videos", Folders: "porn_vid"},
	{ID: "1121676815161901136", Server_ID: "946803206091059232", Server_Name: "Cumsluts_18+", Name: "porn_images", Folders: "porn_img"},
	{ID: "1017082076719218770", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "waifus", Folders: "hentai"},
	{ID: "1017085701805842542", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "futanari", Folders: "hentai,futa"},
	{ID: "1017086918749601842", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "dark", Folders: "hentai"},
	{ID: "1017088174700380200", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "monster", Folders: "hentai"},
	{ID: "1017089863390089267", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "petsgirls", Folders: "hentai"},
	{ID: "1017234431003602985", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "bondage", Folders: "hentai"},
	{ID: "1017235609380720701", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "netorare", Folders: "hentai"},
	{ID: "1017236443304837151", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "hmemes", Folders: "memes"},
	{ID: "1017237893951004682", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "hfavs", Folders: "hentai"},
	{ID: "1017239396698820658", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "milk", Folders: "hentai"},
	{ID: "1017449760841547796", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "unlimited", Folders: "hentai"},
	{ID: "1017451210648854639", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "wallpapers", Folders: "hentai,wallpapers"},
	{ID: "1017453491867897936", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "femdom", Folders: "hentai"},
	{ID: "1017454654277308557", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "cum", Folders: "hentai"},
	{ID: "1017455679134519377", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "blowjob", Folders: "hentai,blowjob"},
	{ID: "1023617222167507124", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "titfuck", Folders: "hentai,titfuck"},
	{ID: "1023618842397786182", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "males", Folders: "hentai"},
	{ID: "1024357950850084895", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "3d", Folders: "hentai"},
	{ID: "1024359555716616312", Server_ID: "1016778742749728798", Server_Name: "NSFW_LOAD", Name: "games", Folders: "hentai"},
	{ID: "630620808930394122", Server_ID: "619487627858411520", Server_Name: "CAJOI_Bunnybutt_Edition", Name: "ecchi_waifus", Folders: "ecchi"},
	{ID: "619487801402064907", Server_ID: "619487627858411520", Server_Name: "CAJOI_Bunnybutt_Edition", Name: "waifu_lewds", Folders: "hentai"},
	{ID: "1060644432816378006", Server_ID: "1060526448722591755", Server_Name: "spyly", Name: "17", Folders: "test,test2"},
	{ID: "935640637036961842", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "ecchi_or_hentai", Folders: "hentai"},
	{ID: "942916127896518666", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "left_or_right", Folders: "hentai"},
	{ID: "942930786271764531", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "pick_a_hole", Folders: "hentai"},
	{ID: "1095752614248976445", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "sub_or_dom", Folders: "hentai"},
	{ID: "943114288732704798", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "suck_or_fuck", Folders: "hentai"},
	{ID: "948259260423352331", Server_ID: "915696409687257149", Server_Name: "Shimatsuri", Name: "yay_or_nay", Folders: "hentai"},
	{ID: "926188746834055198", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Nino", Folders: "hentai"},
	{ID: "926189314612817930", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Miku", Folders: "hentai"},
	{ID: "926189410700120125", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Sakura", Folders: "hentai"},
	{ID: "926192630977667113", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Yotsuba", Folders: "hentai"},
	{ID: "926189382426300478", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Itsuki", Folders: "hentai"},
	{ID: "926194603483340840", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Hanabi", Folders: "hentai"},
	{ID: "926198069043810345", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Konan", Folders: "hentai"},
	{ID: "926198748231639070", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Hayasaka-AI", Folders: "hentai"},
	{ID: "926201553797394442", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Shinobu", Folders: "hentai"},
	{ID: "926202662444204042", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Toga", Folders: "hentai"},
	{ID: "926206442497777736", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "02", Folders: "hentai"},
	{ID: "926206734635241542", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Saber", Folders: "hentai"},
	{ID: "926207123887628360", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Momo", Folders: "hentai"},
	{ID: "926210408551944302", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Ram", Folders: "hentai"},
	{ID: "926216376128000070", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Iris", Folders: "hentai"},
	{ID: "926218422688309358", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Sumi", Folders: "hentai"},
	{ID: "926219205597089802", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Nero", Folders: "hentai"},
	{ID: "926218450949517352", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Erza", Folders: "hentai"},
	{ID: "926222079399964682", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Wallenstein", Folders: "hentai"},
	{ID: "926219172671799306", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Rukia", Folders: "hentai"},
	{ID: "926223192157212702", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Aqua", Folders: "hentai"},
	{ID: "926223956040618014", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Shino", Folders: "hentai"},
	{ID: "926965494827859999", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Albedo", Folders: "hentai"},
	{ID: "926965494827859999", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Android18", Folders: "hentai"},
	{ID: "926966676996628480", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Rias-gremory", Folders: "hentai"},
	{ID: "926968161859928164", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Daki", Folders: "hentai"},
	{ID: "927925371989536818", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Lucy", Folders: "hentai"},
	{ID: "927926239988158475", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "nami", Folders: "hentai"},
	{ID: "927927409234284544", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "mitsuri", Folders: "hentai"},
	{ID: "933595943343902740", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "ayaka", Folders: "hentai"},
	{ID: "942677766665560125", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "tohru", Folders: "hentai"},
	{ID: "961410192770236446", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Tatsumaki", Folders: "hentai"},
	{ID: "961410458483564574", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Nobara", Folders: "hentai"},
	{ID: "962455796862570536", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Kanao", Folders: "hentai"},
	{ID: "962457198460555286", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "TenTen", Folders: "hentai"},
	{ID: "962458681889406976", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Ino", Folders: "hentai"},
	{ID: "964490361630261258", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Sayu", Folders: "hentai"},
	{ID: "964490589519368233", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Uraraka", Folders: "hentai"},
	{ID: "964491125836628018", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Elaina", Folders: "hentai"},
	{ID: "964491285891276820", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "Elizabeth7ds", Folders: "hentai"},
	{ID: "968232497278255134", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "asui-tsuyu", Folders: "hentai"},
	{ID: "968237630326210600", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "raphtailia", Folders: "hentai"},
	{ID: "968238162793078805", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "mount-lady", Folders: "hentai"},
	{ID: "968238383073755176", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "yunyun", Folders: "hentai"},
	{ID: "968238483435028530", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "ruka", Folders: "hentai"},
	{ID: "968238707662544897", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "chizuru-mizuhara", Folders: "hentai"},
	{ID: "968239274031988736", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "kushina", Folders: "hentai"},
	{ID: "968239584435658802", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "uzaki", Folders: "hentai"},
	{ID: "969823673777291334", Server_ID: "926113046772981760", Server_Name: "Hentai_Hub", Name: "ako", Folders: "hentai"},
	{ID: "763570448872112159", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "hentai_entries", Folders: "hentai"},
	{ID: "678663154233901057", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "hentai_gifs", Folders: "hentai"},
	{ID: "773021584213934088", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "hentai_winners", Folders: "hentai"},
	{ID: "736103879661322240", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "just_ecchi", Folders: "ecchi"},
	{ID: "638857197274660865", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "just_hentai", Folders: "hentai"},
	{ID: "655487317741666318", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "seasonal", Folders: "hentai"},
	{ID: "678475258109624330", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "links_and_videos", Folders: "hentai"},
	{ID: "625171628065685524", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "nsfw_bot", Folders: "hentai"},
	{ID: "699851614386651156", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "this_art_style", Folders: "hentai"},
	{ID: "646539287747100686", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "armpit", Folders: "hentai,kinks"},
	{ID: "678406754614247434", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "asshole", Folders: "hentai,kinks"},
	{ID: "638857940635222041", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "bdsm_bondage", Folders: "hentai,kinks"},
	{ID: "646217937186848778", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "breath_play", Folders: "hentai,kinks"},
	{ID: "644403346043699210", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "feet", Folders: "hentai,kinks"},
	{ID: "722881826691088464", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "femdom", Folders: "hentai,kinks"},
	{ID: "638927639570284545", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "futa", Folders: "hentai,kinks"},
	{ID: "646569505811857413", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "lactation", Folders: "hentai,kinks"},
	{ID: "656878824117501962", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "milf", Folders: "hentai,kinks"},
	{ID: "678500910338408458", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "monster_girl", Folders: "hentai,kinks"},
	{ID: "646217312676085760", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "pet_play", Folders: "hentai,kinks"},
	{ID: "638858256525033492", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "tentacle", Folders: "hentai,kinks"},
	{ID: "638871228177842224", Server_ID: "625167654919077889", Server_Name: "why_not_umu", Name: "traps", Folders: "hentai,kinks"},
	{ID: "865683757796818954", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "ecchi", Folders: "ecchi"},
	{ID: "857918054063013888", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "random_hentai", Folders: "hentai"},
	{ID: "857958877094346782", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "3d", Folders: "hentai"},
	{ID: "865644364897452042", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "boobs", Folders: "hentai"},
	{ID: "865644349173399572", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "ass", Folders: "hentai"},
	{ID: "857919804018393098", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "yuri", Folders: "hentai"},
	{ID: "857949246715330570", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "milf", Folders: "hentai"},
	{ID: "1061170949505953832", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "tomboy", Folders: "hentai"},
	{ID: "857949970576834602", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "tanned_skin", Folders: "hentai"},
	{ID: "865637686190800928", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "harem", Folders: "hentai"},
	{ID: "865231951622438932", Server_ID: "857910034725863444", Server_Name: "NSFW_Research", Name: "loli", Folders: "hentai"},
	//{ID: "", Server_ID: "", Server_Name: "", Name: "", Folders: ""},

}

func check_list(ID string) bool {
	there := false
	for i, name := range names {
		if name.ID == ID {
			fmt.Println(i)
			there = true
		}
	}
	return there
}

func download_attachments() {
	for {
		if downloads != nil {
			fmt.Println("Downloading: ", downloads[0])
			parts := strings.SplitN(downloads[0], "*#*", -1)

			URL := parts[0]

			folders := strings.SplitN(parts[5], ",", -1)
			for i := range folders {
				filepath := "data/" + folders[i] + "/" + parts[6]
				fmt.Println("Filepath: ", filepath)

				file, err := os.Create(filepath)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return
				}
				defer file.Close()

				// Download the attachment using its URL.
				resp, err := http.Get(URL)
				if err != nil {
					fmt.Println("Error downloading attachment:", err)
					return
				}
				defer resp.Body.Close()

				// Save the attachment contents to the file.
				_, err = io.Copy(file, resp.Body)
				if err != nil {
					fmt.Println("Error saving attachment:", err)
					return
				}

				fmt.Println("Attachment downloaded:", filepath)
				fmt.Println()

				if len(downloads) > 1 {
					downloads = downloads[1:]
				} else {
					downloads = nil
				}
			}

		}
	}
}

func main() {

	go download_attachments()
	flag.StringVar(&Token, "token", "", "Include the bots token")
	flag.Parse()

	if len(Token) < 10 {
		fmt.Println("Error in Token")
		for {
		}
	}
	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	bot.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	bot.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if check_list(message.ChannelID) != true {
		fmt.Println("Ignoring unlisted channel")
		return
	}

	if len(message.Attachments) == 0 {
		fmt.Println("Ignoring message without attachments")
		return
	}

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == session.State.User.ID {
		fmt.Println("Ignoring own message")
		return
	}

	if len(message.Attachments) != 0 {
		for i, name := range names {
			if name.ID == message.ChannelID {
				fmt.Println(i)
				discord.WebhookURL = "https://discord.com/api/webhooks/1135243087128559687/_zVFIwoQ-1FPAEw7839H6JJlz2aS2TlHMEUg8ikJKe4qfjwBrfmF-yiEAwUOKZog9Hny"
				discord.Say(fmt.Sprint(len(message.Attachments)) + " Attachment(s) found in " + message.ChannelID + "(" + name.Server_Name + ")")

			}
		}

		string1 := fmt.Sprint(len(message.Attachments)) + " Attachment(s) found"
		fmt.Println(string1)
		//session.ChannelMessageSend(message.ChannelID, "Attachment found")

		for _, attachment := range message.Attachments {
			for i, name := range names {
				if name.ID == message.ChannelID {
					fmt.Println(i)
					ID_1 = name.ID
					Server_ID_1 = name.Server_ID
					Server_Name_1 = name.Server_Name
					Name_1 = name.Name
					Folders_1 = name.Folders

				}
			}
			string2 := attachment.URL + "*#*" + ID_1 + "*#*" + Server_ID_1 + "*#*" + Server_Name_1 + "*#*" + Name_1 + "*#*" + Folders_1 + "*#*" + attachment.Filename
			fmt.Println(string2)
			downloads = append(downloads, string2)
			ID_1 = ""
			Server_ID_1 = ""
			Server_Name_1 = ""
			Name_1 = ""
			Folders_1 = ""
		}
	}
}
