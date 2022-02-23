Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.1.1/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/metadata/PutMetadata.json
func ExampleMetadataClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewMetadataClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<metadata-name>",
		armsecurityinsights.MetadataModel{
			Properties: &armsecurityinsights.MetadataProperties{
				Author: &armsecurityinsights.MetadataAuthor{
					Name:  to.StringPtr("<name>"),
					Email: to.StringPtr("<email>"),
				},
				Categories: &armsecurityinsights.MetadataCategories{
					Domains: []*string{
						to.StringPtr("Application"),
						to.StringPtr("Security – Insider Threat")},
					Verticals: []*string{
						to.StringPtr("Healthcare")},
				},
				ContentID: to.StringPtr("<content-id>"),
				Dependencies: &armsecurityinsights.MetadataDependencies{
					Criteria: []*armsecurityinsights.MetadataDependencies{
						{
							Criteria: []*armsecurityinsights.MetadataDependencies{
								{
									Name:      to.StringPtr("<name>"),
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsights.Kind("DataConnector").ToPtr(),
								},
								{
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsights.Kind("DataConnector").ToPtr(),
								},
								{
									ContentID: to.StringPtr("<content-id>"),
									Kind:      armsecurityinsights.Kind("DataConnector").ToPtr(),
									Version:   to.StringPtr("<version>"),
								}},
							Operator: armsecurityinsights.Operator("OR").ToPtr(),
						},
						{
							ContentID: to.StringPtr("<content-id>"),
							Kind:      armsecurityinsights.Kind("Playbook").ToPtr(),
							Version:   to.StringPtr("<version>"),
						},
						{
							ContentID: to.StringPtr("<content-id>"),
							Kind:      armsecurityinsights.Kind("Parser").ToPtr(),
						}},
					Operator: armsecurityinsights.Operator("AND").ToPtr(),
				},
				FirstPublishDate: to.TimePtr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				Kind:             armsecurityinsights.Kind("AnalyticsRule").ToPtr(),
				LastPublishDate:  to.TimePtr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
				ParentID:         to.StringPtr("<parent-id>"),
				Providers: []*string{
					to.StringPtr("Amazon"),
					to.StringPtr("Microsoft")},
				Source: &armsecurityinsights.MetadataSource{
					Name:     to.StringPtr("<name>"),
					Kind:     armsecurityinsights.SourceKind("Solution").ToPtr(),
					SourceID: to.StringPtr("<source-id>"),
				},
				Support: &armsecurityinsights.MetadataSupport{
					Name:  to.StringPtr("<name>"),
					Email: to.StringPtr("<email>"),
					Link:  to.StringPtr("<link>"),
					Tier:  armsecurityinsights.SupportTier("Partner").ToPtr(),
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
