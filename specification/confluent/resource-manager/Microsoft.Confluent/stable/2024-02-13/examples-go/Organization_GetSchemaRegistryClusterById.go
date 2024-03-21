package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_GetSchemaRegistryClusterById.json
func ExampleOrganizationClient_GetSchemaRegistryClusterByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().GetSchemaRegistryClusterByID(ctx, "myResourceGroup", "myOrganization", "env-stgcczjp2j3", "lsrc-stgczkq22z", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SchemaRegistryClusterRecord = armconfluent.SchemaRegistryClusterRecord{
	// 	ID: to.Ptr("lsrc-stgczkq22z"),
	// 	Kind: to.Ptr("Cluster"),
	// 	Properties: &armconfluent.SchemaRegistryClusterProperties{
	// 		Metadata: &armconfluent.SCMetadataEntity{
	// 			CreatedTimestamp: to.Ptr("2023-11-08T07:37:14.309386Z"),
	// 			DeletedTimestamp: to.Ptr("2023-11-08T07:37:14.309386Z"),
	// 			ResourceName: to.Ptr("crn://confluent.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/environment=env-stgccnk2mgd/schema-registry=lsrc-stgczkq22z"),
	// 			Self: to.Ptr("https://api.stag.cpdev.cloud/srcm/v2/clusters/lsrc-stgczkq22z"),
	// 			UpdatedTimestamp: to.Ptr("2023-11-08T07:37:14.309386Z"),
	// 		},
	// 		Spec: &armconfluent.SchemaRegistryClusterSpecEntity{
	// 			Name: to.Ptr("Stream Governance Package"),
	// 			Cloud: to.Ptr("GCP"),
	// 			Environment: &armconfluent.SchemaRegistryClusterEnvironmentRegionEntity{
	// 				ID: to.Ptr("env-stgccnk2mgd"),
	// 				Related: to.Ptr("https://api.stag.cpdev.cloud/org/v2/environments/env-stgccnk2mgd"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=37d1220b-93a6-43e3-a114-dd8a20a94a31/environment=env-stgccnk2mgd"),
	// 			},
	// 			HTTPEndpoint: to.Ptr("https://psrc-57wzyg.centralus.azure.stag.cpdev.cloud"),
	// 			Package: to.Ptr("ADVANCED"),
	// 			Region: &armconfluent.SchemaRegistryClusterEnvironmentRegionEntity{
	// 				ID: to.Ptr("sgreg-7"),
	// 				Related: to.Ptr("https://api.stag.cpdev.cloud/srcm/v2/regions/sgreg-7"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/schema-registry-region=sgreg-7"),
	// 			},
	// 		},
	// 		Status: &armconfluent.SchemaRegistryClusterStatusEntity{
	// 			Phase: to.Ptr("PROVISIONED"),
	// 		},
	// 	},
	// }
}
