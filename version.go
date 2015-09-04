package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Version struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	Required    bool   `json:"required"`
}

type Versions []Version

func (vs Versions) Next(curr string) (string, error) {
	found := false
	nextRequired := ""
	nextPublished := ""

	for _, v := range vs {
		if v.Version == curr {
			found = true
			continue
		}

		if found && v.Published {
			nextPublished = v.Version

			if v.Required {
				nextRequired = v.Version
				break
			}
		}
	}

	if !found {
		return "", fmt.Errorf("current version %q not found", curr)
	}

	if nextRequired != "" {
		return nextRequired, nil
	}

	if nextPublished != "" {
		return nextPublished, nil
	}

	return "", fmt.Errorf("current version %q is latest", curr)
}

func (vs Versions) Latest() (string, error) {
	for i := len(vs) - 1; i >= 0; i-- {
		v := vs[i]

		if v.Published {
			return v.Version, nil
		}
	}

	return "", fmt.Errorf("no published versions")
}

// Append a new version to versions.json file
func AppendVersion(v Version) (*Versions, error) {
	vs, err := getVersions()

	if err != nil {
		return nil, err
	}

	*vs = append(*vs, v)

	err = putVersions(*vs)

	return vs, nil
}

// Walk a bucket to create initial versions.json file
func ImportVersions() (*Versions, error) {
	vs := Versions{}

	S3 := s3.New(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	})

	res, err := S3.ListObjects(&s3.ListObjectsInput{
		Bucket:    aws.String("convox"),
		Delimiter: aws.String("/"),
		Prefix:    aws.String("release/"),
	})

	if err != nil {
		return nil, err
	}

	for _, p := range res.CommonPrefixes {
		parts := strings.Split(*p.Prefix, "/")
		version := parts[1]

		if version == "latest" {
			continue
		}

		vs = append(vs, Version{
			Version:   version,
			Published: false,
			Required:  false,
		})
	}

	err = putVersions(vs)

	return &vs, err
}

func getVersions() (*Versions, error) {
	vs := Versions{}

	S3 := s3.New(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	})

	res, err := S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("convox"),
		Key:    aws.String("release/versions.json"),
	})

	if err != nil && err.(awserr.Error).Code() != "NoSuchKey" {
		return nil, err
	} else {
		b, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		json.Unmarshal(b, &vs)
	}

	return &vs, nil
}

func putVersions(vs Versions) error {
	data, err := json.MarshalIndent(vs, "", "  ")

	if err != nil {
		return err
	}

	S3 := s3.New(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	})

	_, err = S3.PutObject(&s3.PutObjectInput{
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(data),
		Bucket:        aws.String("convox"),
		ContentLength: aws.Int64(int64(len(data))),
		Key:           aws.String("release/versions.json"),
	})

	return err
}
