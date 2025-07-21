package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/ListAuthKeysSqlMigrationService.json
func ExampleSQLMigrationServicesClient_ListAuthKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLMigrationServicesClient().ListAuthKeys(ctx, "testrg", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthenticationKeys = armdatamigration.AuthenticationKeys{
	// 	AuthKey1: to.Ptr("IR@3c400518-8dce-479f-97f9-83d17c3d8c24@dmsv2test7@ServiceEndpoint=wu.frontend.int.clouddatahub-int.net@Kjvm37BceBapP0z7R2n3zZrvY/M79yo7Hg3rR+LjFFI="),
	// 	AuthKey2: to.Ptr("IR@3c499528-8ace-47xf-97z9-83d17c3d8c24@dmsv2test7@ServiceEndpoint=wu.frontend.int.clouddatahub-int.net@Kjvm37BceBapP0z7R2n3zZrvY/M79yo7Hg3rR+LjFFI="),
	// }
}
