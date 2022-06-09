```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/metadata/PutMetadata.json
func ExampleMetadataClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewMetadataClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<metadata-name>",
		armsecurityinsight.MetadataModel{
			Properties: &armsecurityinsight.MetadataProperties{
				Author: &armsecurityinsight.MetadataAuthor{
					Name:  to.StringPtr("<name>"),
					Email: to.StringPtr("<email>"),
				},
				Categories: &armsecurityinsight.MetadataCategories{
					Domains: []*string{
						to.StringPtr("Application"),
						to.StringPtr("Security – Insider Threat")},
					Verticals: []*string{
						to.StringPtr("Healthcare")},
				},
				ContentID: to.StringPtr("<content-id>"),
				Dependencies: &armsecurityinsight.MetadataDependencies{
					Criteria: []*armsecurityinsight.MetadataDependencies{
						{
							Criteria: []*armsecurityinsight.MetadataDependencies{
								{
									Name:      to.StringPtr("<name>"),
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsight.Kind("DataConnector").ToPtr(),
								},
								{
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsight.Kind("DataConnector").ToPtr(),
								},
								{
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsight.Kind("DataConnector").ToPtr(),
									Version:   to.StringPtr("<version>"),
								}},
							Operator: armsecurityinsight.Operator("OR").ToPtr(),
						},
						{
							ContentID: to.StringPtr("<content-id>"),
							Kind:      armsecurityinsight.Kind("Playbook").ToPtr(),
							Version:   to.StringPtr("<version>"),
						},
						{
							ContentID: to.StringPtr("<content-id>"),
							Kind:      armsecurityinsight.Kind("Parser").ToPtr(),
						}},
					Operator: armsecurityinsight.Operator("AND").ToPtr(),
				},
				FirstPublishDate: to.TimePtr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				Kind:             armsecurityinsight.Kind("AnalyticsRule").ToPtr(),
				LastPublishDate:  to.TimePtr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				ParentID:         to.StringPtr("<parent-id>"),
				Providers: []*string{
					to.StringPtr("Amazon"),
					to.StringPtr("Microsoft")},
				Source: &armsecurityinsight.MetadataSource{
					Name:     to.StringPtr("<name>"),
					Kind:     armsecurityinsight.SourceKind("Solution").ToPtr(),
					SourceID: to.StringPtr("<source-id>"),
				},
				Support: &armsecurityinsight.MetadataSupport{
					Name:  to.StringPtr("<name>"),
					Email: to.StringPtr("<email>"),
					Link:  to.StringPtr("<link>"),
					Tier:  armsecurityinsight.SupportTier("Partner").ToPtr(),
				},
				Version: to.StringPtr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MetadataClientCreateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.
