//go:generate mockgen -destination=mock/mock_opensearch_client.go -package=mock github.com/krobus00/krokit OpensearchClient

package kit

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/goccy/go-json"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
)

type OSConfig struct {
	Addresses          []string
	InsecureSkipVerify bool
	Username           string
	Password           string
}

type IndexModel interface {
	GetID() string
}

type OpensearchClient interface {
	Index(ctx context.Context, indexName string, model IndexModel) error
	CreateIndices(ctx context.Context, indexName string, body *strings.Reader) (*opensearchapi.Response, error)
	PutIndicesMapping(ctx context.Context, indexNames []string, body *strings.Reader) (*opensearchapi.Response, error)
}

type opensearchClient struct {
	client *opensearch.Client
}

func NewOpensearchClient(config *OSConfig) (OpensearchClient, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.InsecureSkipVerify},
		},
		Addresses: config.Addresses,
		Username:  config.Username,
		Password:  config.Password,
	})
	krokitOSClient := &opensearchClient{
		client: client,
	}
	return krokitOSClient, err
}

func (k *opensearchClient) CreateIndices(ctx context.Context, indexName string, body *strings.Reader) (*opensearchapi.Response, error) {
	logger := logrus.WithFields(logrus.Fields{
		"indexName": indexName,
	})
	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  body,
	}

	res, err := req.Do(ctx, k.client)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	logger.Info(res)

	return res, err
}

func (k *opensearchClient) PutIndicesMapping(ctx context.Context, indexNames []string, body *strings.Reader) (*opensearchapi.Response, error) {
	logger := logrus.WithFields(logrus.Fields{
		"indexNames": indexNames,
	})

	req := opensearchapi.IndicesPutMappingRequest{
		Index: indexNames,
		Body:  body,
	}

	res, err := req.Do(ctx, k.client)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	logger.Info(res)

	return res, nil

}

func (k *opensearchClient) Index(ctx context.Context, indexName string, model IndexModel) error {
	logger := logrus.WithFields(logrus.Fields{
		"indexName": indexName,
		"docID":     model.GetID(),
	})

	docData, err := json.Marshal(model)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	body := strings.NewReader(string(docData))

	req := opensearchapi.IndexRequest{
		Index:      indexName,
		DocumentID: model.GetID(),
		Body:       body,
		Pretty:     true,
	}

	res, err := req.Do(ctx, k.client)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer res.Body.Close()

	logger.Info(res)

	return nil
}
