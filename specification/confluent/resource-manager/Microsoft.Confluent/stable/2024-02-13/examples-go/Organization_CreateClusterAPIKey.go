package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_CreateClusterAPIKey.json
func ExampleOrganizationClient_CreateAPIKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().CreateAPIKey(ctx, "myResourceGroup", "myOrganization", "env-12132", "clusterId-123", armconfluent.CreateAPIKeyModel{
		Name:        to.Ptr("CI kafka access key"),
		Description: to.Ptr("This API key provides kafka access to cluster x"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeyRecord = armconfluent.APIKeyRecord{
	// 	ID: to.Ptr("JDCXTFUVFVQGF56J"),
	// 	Kind: to.Ptr("ApiKey"),
	// 	Properties: &armconfluent.APIKeyProperties{
	// 		Metadata: &armconfluent.SCMetadataEntity{
	// 			CreatedTimestamp: to.Ptr("2023-11-15T16:15:44.489498Z"),
	// 			ResourceName: to.Ptr("crn://api.stag.cpdev.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/service-account=sa-wwn7mm/api-key=JDCXTFUVFVQGF56J"),
	// 			Self: to.Ptr("https://api.stag.cpdev.cloud/iam/v2/api-keys/JDCXTFUVFVQGF56J"),
	// 			UpdatedTimestamp: to.Ptr("2023-11-15T16:15:44.489498Z"),
	// 		},
	// 		Spec: &armconfluent.APIKeySpecEntity{
	// 			Name: to.Ptr("CI kafka access key"),
	// 			Description: to.Ptr("This API key provides kafka access to cluster x"),
	// 			Owner: &armconfluent.APIKeyOwnerEntity{
	// 				ID: to.Ptr("sa-wwn7mm"),
	// 				Kind: to.Ptr("ServiceAccount"),
	// 				Related: to.Ptr("https://api.stag.cpdev.cloud/iam/v2/service-accounts/sa-wwn7mm"),
	// 				ResourceName: to.Ptr("crn://api.stag.cpdev.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/service-account=sa-wwn7mm"),
	// 			},
	// 			Resource: &armconfluent.APIKeyResourceEntity{
	// 				ID: to.Ptr("lsrc-stgc1yrzz3"),
	// 				Kind: to.Ptr("SchemaRegistry"),
	// 				Related: to.Ptr("https://api.stag.cpdev.cloud/srcm/v2/schema-registries/lsrc-stgc1yrzz3"),
	// 				ResourceName: to.Ptr("crn://api.stag.cpdev.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/schema-registry=lsrc-stgc1yrzz3"),
	// 			},
	// 			Secret: to.Ptr(""),
	// 		},
	// 	},
	// }
}
