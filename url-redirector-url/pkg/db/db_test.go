package db_test

import (
	"testing"
	"url-redirecter-url/pkg/models"
	"url-redirecter-url/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestInsertURL(t *testing.T) {
	insertURL(t)
}

func TestGetActiveURL(t *testing.T) {
	// rand1 := randomUrl()

	// active1, err := TestStorage.InsertURL(rand1)

	ins1 := insertURL(t)

	userId := ins1.UserID

	_, err := TestStorage.SetActive(ins1.ID)

	res, err := TestStorage.GetActiveURL(userId)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.True(t, res.Active)
}

func TestSetActive(t *testing.T) {
	url := insertURL(t)
	urlId, err := TestStorage.SetActive(url.ID)
	require.NoError(t, err)
	resUrl, err := TestStorage.GetURL(urlId)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.True(t, resUrl.Active)
}

func TestSetNotActive(t *testing.T) {
	url := insertURL(t)
	activeUrlId, err := TestStorage.SetActive(url.ID)
	require.NoError(t, err)

	notActiveUrlId, err := TestStorage.SetNotActive(activeUrlId)
	require.NoError(t, err)

	resUrl, err := TestStorage.GetURL(notActiveUrlId)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.False(t, resUrl.Active)
}

func TestGetURL(t *testing.T) {
	url := insertURL(t)

	res, err := TestStorage.GetURL(url.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, url, res)
}

func TestGetUserURLs(t *testing.T) {
	userID := utils.RandomInt(1, 1000)
	var urls []models.URL
	for i := 0; i < 10; i++ {
		url := models.URL{
			UserID: userID,
			URL:    utils.RandomString(10),
			Active: false,
		}

		resUrl, err := TestStorage.InsertURL(url)
		require.NoError(t, err)

		urls = append(urls, resUrl)
	}

	require.Equal(t, len(urls), 10)
	for i := 0; i < 10; i++ {
		require.Equal(t, urls[i].UserID, userID)
	}
}

func randomUrl() models.URL {
	return models.URL{
		UserID: utils.RandomInt(1, 50),
		URL:    utils.RandomString(10),
		Active: false,
	}
}

func insertURL(t *testing.T) models.URL {
	randUrl := randomUrl()
	resUrl, err := TestStorage.InsertURL(randUrl)
	require.NoError(t, err)
	require.NotEmpty(t, resUrl)
	require.Equal(t, randUrl.UserID, resUrl.UserID)
	require.Equal(t, randUrl.URL, resUrl.URL)
	return resUrl
}
