package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_GetClusterAPIKey.json
func ExampleOrganizationClient_GetClusterAPIKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().GetClusterAPIKey(ctx, "myResourceGroup", "myOrganization", "apiKeyId-123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeyRecord = armconfluent.APIKeyRecord{
	// 	ID: to.Ptr("apiKeyId-123"),
	// 	Kind: to.Ptr("ApiKey"),
	// 	Properties: &armconfluent.APIKeyProperties{
	// 		Metadata: &armconfluent.SCMetadataEntity{
	// 			CreatedTimestamp: to.Ptr("2023-11-15T16:15:44.489498Z"),
	// 			ResourceName: to.Ptr("crn://api.stag.cpdev.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/service-account=sa-wwn7mm/api-key=apiKeyId-123"),
	// 			Self: to.Ptr("https://api.stag.cpdev.cloud/iam/v2/api-keys/apiKeyId-123"),
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
	// 				Environment: to.Ptr("env-0000"),
	// 				ID: to.Ptr("lsrc-stgc1yrzz3"),
	// 				Kind: to.Ptr("SchemaRegistry"),
	// 				Related: to.Ptr("https://api.stag.cpdev.cloud/srcm/v2/schema-registries/lsrc-stgc1yrzz3"),
	// 				ResourceName: to.Ptr("crn://api.stag.cpdev.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/schema-registry=lsrc-stgc1yrzz3"),
	// 			},
	// 		},
	// 	},
	// }
}
