package services

import (
	"fmt"
	"io"
	"net/http"

	. "github.com/fagom/gitc/internal"
	. "github.com/fagom/gitc/models"
)

func CheckUserAuthKey(user User) error {
	client := &http.Client{}
	addr := ParseUrl(GithubRootPath, GithubUserPath)
	token := "Bearer " + user.PassKey

	req, err := http.NewRequest("GET", addr, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))

	return nil
}
