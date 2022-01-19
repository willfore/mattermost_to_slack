package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/willfore/mattermost_to_slack/types"
)

func FetchUsers() (types.SlackUserList, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://slack.com/api/users.list?limit=100&cursor=&include_locale=", nil)
	if err != nil {
		fmt.Errorf("could not create request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SLACK_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("could not send request: %s", err)
	}

	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Errorf("could not read response body: %s", err)
	//}

	//bodyString := string(bodyBytes)

	var slackUserList types.SlackUserList
	json.NewDecoder(resp.Body).Decode(&slackUserList)
	if err != nil {
		fmt.Errorf("could not decode response: %s", err)
	}

	defer resp.Body.Close()
	return slackUserList, nil
}
