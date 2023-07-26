package db_test

import (
	"testing"

	"url-redirecter-url/pkg/models"
	"url-redirecter-url/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestInsertUrl(t *testing.T) {
	insertUrl(t)
}

func TestGetActiveUrl(t *testing.T) {
	// rand1 := randomUrl()

	// active1, err := TestStorage.InsertUrl(rand1)

	ins1 := insertUrl(t)

	userId := ins1.UserID

	_, err := TestStorage.Activate(ins1.ID)

	res, err := TestStorage.GetActiveUrl(userId)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.True(t, res.Active)
}

func TestSetActive(t *testing.T) {
	url := insertUrl(t)
	urlId, err := TestStorage.Activate(url.ID)
	require.NoError(t, err)
	resUrl, err := TestStorage.GetUrl(urlId)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.True(t, resUrl.Active)
}

func TestSetNotActive(t *testing.T) {
	url := insertUrl(t)
	activeUrlId, err := TestStorage.Activate(url.ID)
	require.NoError(t, err)

	notActiveUrlId, err := TestStorage.Deactivate(activeUrlId)
	require.NoError(t, err)

	resUrl, err := TestStorage.GetUrl(notActiveUrlId)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.False(t, resUrl.Active)
}

func TestGetUrl(t *testing.T) {
	url := insertUrl(t)

	res, err := TestStorage.GetUrl(url.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, url, res)
}

func TestGetUserUrls(t *testing.T) {
	userID := utils.RandomInt(1, 1000)
	var urls []models.Url
	for i := 0; i < 10; i++ {
		url := models.Url{
			UserID: userID,
			Url:    utils.RandomString(10),
			Active: false,
		}

		resUrl, err := TestStorage.InsertUrl(url)
		require.NoError(t, err)

		urls = append(urls, resUrl)
	}

	require.Equal(t, len(urls), 10)
	for i := 0; i < 10; i++ {
		require.Equal(t, urls[i].UserID, userID)
	}
}

func randomUrl() models.Url {
	return models.Url{
		UserID: utils.RandomInt(1, 50),
		Url:    utils.RandomString(10),
		Active: false,
	}
}

func insertUrl(t *testing.T) models.Url {
	randUrl := randomUrl()
	resUrl, err := TestStorage.InsertUrl(randUrl)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.Equal(t, randUrl.UserID, resUrl.UserID)
	require.Equal(t, randUrl.Url, resUrl.Url)
	return resUrl
}
