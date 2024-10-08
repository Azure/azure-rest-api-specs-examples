package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/56537883b7cdb95618c3d1ec1c0ee37b59d88d72/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CheckDependencies_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_CheckDependencies_serverlessRuntimesCheckDependencies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().CheckDependencies(ctx, "rgopenapi", "3P", "M", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckDependenciesResponse = arminformaticadatamgmt.CheckDependenciesResponse{
	// 	Count: to.Ptr[int32](28),
	// 	ID: to.Ptr("uwjsqpxr"),
	// 	References: []*arminformaticadatamgmt.ServerlessRuntimeDependency{
	// 		{
	// 			Path: to.Ptr("yxbpmcmfhhtht"),
	// 			Description: to.Ptr("vlkyqkevlrpge"),
	// 			AppContextID: to.Ptr("t"),
	// 			DocumentType: to.Ptr("jpcz"),
	// 			ID: to.Ptr("uzp"),
	// 			LastUpdatedTime: to.Ptr("yyf"),
	// 	}},
	// }
}
