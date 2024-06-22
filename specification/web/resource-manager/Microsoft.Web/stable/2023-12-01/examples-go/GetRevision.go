package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetRevision.json
func ExampleContainerAppsRevisionsClient_GetRevision() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsRevisionsClient().GetRevision(ctx, "rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Revision = armappservice.Revision{
	// 	Name: to.Ptr("testcontainerApp0-pjxhsye"),
	// 	Type: to.Ptr("Microsoft.Web/containerApps/revisions"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.WebcontainerApps/testcontainerApp0/revisions/testcontainerApp0-pjxhsye"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.RevisionProperties{
	// 		Active: to.Ptr(true),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T21:24:22.000Z"); return t}()),
	// 		Fqdn: to.Ptr("testcontainerApp0-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
	// 		Replicas: to.Ptr[int32](1),
	// 		Template: &armappservice.Template{
	// 			Containers: []*armappservice.Container{
	// 				{
	// 					Name: to.Ptr("testcontainerApp0"),
	// 					Image: to.Ptr("repo/testcontainerApp0:v2"),
	// 					Resources: &armappservice.ContainerResources{
	// 						CPU: to.Ptr[float64](0.2),
	// 						Memory: to.Ptr("100Mi"),
	// 					},
	// 			}},
	// 			Dapr: &armappservice.Dapr{
	// 				AppPort: to.Ptr[int32](3000),
	// 				Enabled: to.Ptr(true),
	// 			},
	// 			Scale: &armappservice.Scale{
	// 				MaxReplicas: to.Ptr[int32](5),
	// 				MinReplicas: to.Ptr[int32](1),
	// 				Rules: []*armappservice.ScaleRule{
	// 					{
	// 						Name: to.Ptr("httpscalingrule"),
	// 						HTTP: &armappservice.HTTPScaleRule{
	// 							Metadata: map[string]*string{
	// 								"concurrentRequests": to.Ptr("50"),
	// 							},
	// 						},
	// 				}},
	// 			},
	// 		},
	// 	},
	// }
}
