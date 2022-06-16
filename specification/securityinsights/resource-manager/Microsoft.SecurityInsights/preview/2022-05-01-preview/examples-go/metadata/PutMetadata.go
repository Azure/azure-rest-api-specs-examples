package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/metadata/PutMetadata.json
func ExampleMetadataClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewMetadataClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"myRg",
		"myWorkspace",
		"metadataName",
		armsecurityinsights.MetadataModel{
			Properties: &armsecurityinsights.MetadataProperties{
				Author: &armsecurityinsights.MetadataAuthor{
					Name:  to.Ptr("User Name"),
					Email: to.Ptr("email@microsoft.com"),
				},
				Categories: &armsecurityinsights.MetadataCategories{
					Domains: []*string{
						to.Ptr("Application"),
						to.Ptr("Security – Insider Threat")},
					Verticals: []*string{
						to.Ptr("Healthcare")},
				},
				ContentID:            to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
				ContentSchemaVersion: to.Ptr("2.0"),
				CustomVersion:        to.Ptr("1.0"),
				Dependencies: &armsecurityinsights.MetadataDependencies{
					Criteria: []*armsecurityinsights.MetadataDependencies{
						{
							Criteria: []*armsecurityinsights.MetadataDependencies{
								{
									Name:      to.Ptr("Microsoft Defender for Endpoint"),
									ContentID: to.Ptr("045d06d0-ee72-4794-aba4-cf5646e4c756"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
								},
								{
									ContentID: to.Ptr("dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
								},
								{
									ContentID: to.Ptr("de4dca9b-eb37-47d6-a56f-b8b06b261593"),
									Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
									Version:   to.Ptr("2.0"),
								}},
							Operator: to.Ptr(armsecurityinsights.OperatorOR),
						},
						{
							ContentID: to.Ptr("31ee11cc-9989-4de8-b176-5e0ef5c4dbab"),
							Kind:      to.Ptr(armsecurityinsights.KindPlaybook),
							Version:   to.Ptr("1.0"),
						},
						{
							ContentID: to.Ptr("21ba424a-9438-4444-953a-7059539a7a1b"),
							Kind:      to.Ptr(armsecurityinsights.KindParser),
						}},
					Operator: to.Ptr(armsecurityinsights.OperatorAND),
				},
				FirstPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				Kind:             to.Ptr(armsecurityinsights.KindAnalyticsRule),
				LastPublishDate:  to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				ParentID:         to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
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
					Name:     to.Ptr("Contoso Solution 1.0"),
					Kind:     to.Ptr(armsecurityinsights.SourceKindSolution),
					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
				},
				Support: &armsecurityinsights.MetadataSupport{
					Name:  to.Ptr("Microsoft"),
					Email: to.Ptr("support@microsoft.com"),
					Link:  to.Ptr("https://support.microsoft.com/"),
					Tier:  to.Ptr(armsecurityinsights.SupportTierPartner),
				},
				ThreatAnalysisTactics: []*string{
					to.Ptr("reconnaissance"),
					to.Ptr("commandandcontrol")},
				ThreatAnalysisTechniques: []*string{
					to.Ptr("T1548"),
					to.Ptr("T1548.001")},
				Version: to.Ptr("1.0.0.0"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
