package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/metadata/PutMetadata.json
func ExampleMetadataClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewMetadataClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<metadata-name>",
		armsecurityinsights.MetadataModel{
			Properties: &armsecurityinsights.MetadataProperties{
				Author: &armsecurityinsights.MetadataAuthor{
					Name:  to.Ptr("<name>"),
					Email: to.Ptr("<email>"),
				},
				Categories: &armsecurityinsights.MetadataCategories{
					Domains: []*string{
						to.Ptr("Application"),
						to.Ptr("Security – Insider Threat")},
					Verticals: []*string{
						to.Ptr("Healthcare")},
				},
				ContentID:            to.Ptr("<content-id>"),
				ContentSchemaVersion: to.Ptr("<content-schema-version>"),
				CustomVersion:        to.Ptr("<custom-version>"),
				Dependencies: &armsecurityinsights.MetadataDependencies{
					Criteria: []*armsecurityinsights.MetadataDependencies{
						{
							Criteria: []*armsecurityinsights.MetadataDependencies{
								{
									Name:      to.Ptr("<name>"),
									ContentID: to.Ptr("<content-id>"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
								},
								{
									ContentID: to.Ptr("<content-id>"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
								},
								{
									ContentID: to.Ptr("<content-id>"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
									Version:   to.Ptr("<version>"),
								}},
							Operator: to.Ptr(armsecurityinsights.OperatorOR),
						},
						{
							ContentID: to.Ptr("<content-id>"),
							Kind:      to.Ptr(armsecurityinsights.KindPlaybook),
							Version:   to.Ptr("<version>"),
						},
						{
							ContentID: to.Ptr("<content-id>"),
							Kind:      to.Ptr(armsecurityinsights.KindParser),
						}},
					Operator: to.Ptr(armsecurityinsights.OperatorAND),
				},
				FirstPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				Kind:             to.Ptr(armsecurityinsights.KindAnalyticsRule),
				LastPublishDate:  to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				ParentID:         to.Ptr("<parent-id>"),
				PreviewImages: []*string{
					to.Ptr("firstImage.png"),
					to.Ptr("secondImage.jpeg")},
				PreviewImagesDark: []*string{
					to.Ptr("firstImageDark.png"),
					to.Ptr("secondImageDark.jpeg")},
				Providers: []*string{
					to.Ptr("Amazon"),
					to.Ptr("Microsoft")},
				Source: &armsecurityinsights.MetadataSource{
					Name:     to.Ptr("<name>"),
					Kind:     to.Ptr(armsecurityinsights.SourceKindSolution),
					SourceID: to.Ptr("<source-id>"),
				},
				Support: &armsecurityinsights.MetadataSupport{
					Name:  to.Ptr("<name>"),
					Email: to.Ptr("<email>"),
					Link:  to.Ptr("<link>"),
					Tier:  to.Ptr(armsecurityinsights.SupportTierPartner),
				},
				ThreatAnalysisTactics: []*string{
					to.Ptr("reconnaissance"),
					to.Ptr("commandandcontrol")},
				ThreatAnalysisTechniques: []*string{
					to.Ptr("T1548"),
					to.Ptr("T1548.001")},
				Version: to.Ptr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
