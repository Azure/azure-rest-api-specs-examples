package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/RestoreWebAppBackup.json
func ExampleWebAppsClient_BeginRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppsClient().BeginRestore(ctx, "testrg123", "sitef6141", "123244", armappservice.RestoreRequest{
		Properties: &armappservice.RestoreRequestProperties{
			Databases: []*armappservice.DatabaseBackupSetting{
				{
					Name:                 to.Ptr("backenddb"),
					ConnectionString:     to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
					ConnectionStringName: to.Ptr("backend"),
					DatabaseType:         to.Ptr(armappservice.DatabaseTypeSQLAzure),
				},
				{
					Name:                 to.Ptr("statsdb"),
					ConnectionString:     to.Ptr("DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
					ConnectionStringName: to.Ptr("stats"),
					DatabaseType:         to.Ptr(armappservice.DatabaseTypeSQLAzure),
				}},
			Overwrite:         to.Ptr(true),
			SiteName:          to.Ptr("sitef6141"),
			StorageAccountURL: to.Ptr("DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
