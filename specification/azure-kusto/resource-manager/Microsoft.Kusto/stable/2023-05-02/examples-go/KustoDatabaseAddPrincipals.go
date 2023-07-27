package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabaseAddPrincipals.json
func ExampleDatabasesClient_AddPrincipals() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().AddPrincipals(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DatabasePrincipalListRequest{
		Value: []*armkusto.DatabasePrincipal{
			{
				Name:  to.Ptr("Some User"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeUser),
				AppID: to.Ptr(""),
				Email: to.Ptr("user@microsoft.com"),
				Fqn:   to.Ptr("aaduser=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			},
			{
				Name:  to.Ptr("Kusto"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeGroup),
				AppID: to.Ptr(""),
				Email: to.Ptr("kusto@microsoft.com"),
				Fqn:   to.Ptr("aadgroup=some_guid"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleViewer),
			},
			{
				Name:  to.Ptr("SomeApp"),
				Type:  to.Ptr(armkusto.DatabasePrincipalTypeApp),
				AppID: to.Ptr("some_guid_app_id"),
				Email: to.Ptr(""),
				Fqn:   to.Ptr("aadapp=some_guid_app_id"),
				Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabasePrincipalListResult = armkusto.DatabasePrincipalListResult{
	// 	Value: []*armkusto.DatabasePrincipal{
	// 		{
	// 			Name: to.Ptr("Some User"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeUser),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("user@microsoft.com"),
	// 			Fqn: to.Ptr("aaduser=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 		},
	// 		{
	// 			Name: to.Ptr("Kusto"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeGroup),
	// 			AppID: to.Ptr(""),
	// 			Email: to.Ptr("kusto@microsoft.com"),
	// 			Fqn: to.Ptr("aadgroup=some_guid"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleViewer),
	// 		},
	// 		{
	// 			Name: to.Ptr("SomeApp"),
	// 			Type: to.Ptr(armkusto.DatabasePrincipalTypeApp),
	// 			AppID: to.Ptr("some_guid_app_id"),
	// 			Email: to.Ptr(""),
	// 			Fqn: to.Ptr("aadapp=some_guid_app_id"),
	// 			Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 	}},
	// }
}
