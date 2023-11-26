package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/GetSqlServerInstance.json
func ExampleSQLServerInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLServerInstancesClient().Get(ctx, "testrg", "testsqlServerInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLServerInstance = armazurearcdata.SQLServerInstance{
	// 	Name: to.Ptr("testsqlServerInstance"),
	// 	Type: to.Ptr("Microsoft.AzureArcData/SqlServerInstances"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/testsqlServerInstance"),
	// 	SystemData: &armazurearcdata.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"mytag": to.Ptr("myval"),
	// 	},
	// 	Properties: &armazurearcdata.SQLServerInstanceProperties{
	// 		AzureDefenderStatus: to.Ptr(armazurearcdata.DefenderStatusProtected),
	// 		AzureDefenderStatusLastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		Collation: to.Ptr("collation"),
	// 		ContainerResourceID: to.Ptr("Resource id of hosting Arc Machine"),
	// 		CreateTime: to.Ptr("01/01/2020 01:01:01"),
	// 		CurrentVersion: to.Ptr("2012"),
	// 		Edition: to.Ptr(armazurearcdata.EditionTypeDeveloper),
	// 		HostType: to.Ptr(armazurearcdata.HostTypePhysicalServer),
	// 		InstanceName: to.Ptr("name of instance"),
	// 		LicenseType: to.Ptr(armazurearcdata.ArcSQLServerLicenseTypeFree),
	// 		PatchLevel: to.Ptr("patchlevel"),
	// 		ProductID: to.Ptr("sql id"),
	// 		Status: to.Ptr(armazurearcdata.ConnectionStatusRegistered),
	// 		TCPDynamicPorts: to.Ptr("1433"),
	// 		TCPStaticPorts: to.Ptr("1433"),
	// 		VCore: to.Ptr("4"),
	// 		Version: to.Ptr(armazurearcdata.SQLVersionSQLServer2012),
	// 	},
	// }
}
