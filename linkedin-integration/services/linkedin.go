package services

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "linkedin-integration/config"
    "linkedin-integration/models"
)

type LinkedInService interface {
    PublishPost(post models.Post) error
}

type linkedinServiceImpl struct {
    AccessToken string
}

func NewLinkedInService() LinkedInService {
    cfg := config.LoadConfig()
    return &linkedinServiceImpl{AccessToken: cfg.LinkedInAccessToken}
}

func (s *linkedinServiceImpl) PublishPost(post models.Post) error {
    url := "https://api.linkedin.com/v2/shares"
    requestBody, _ := json.Marshal(map[string]string{
        "content": post.Content,
        "title":   post.Title,
    })
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+s.AccessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to publish post")
    }
    return nil
}
