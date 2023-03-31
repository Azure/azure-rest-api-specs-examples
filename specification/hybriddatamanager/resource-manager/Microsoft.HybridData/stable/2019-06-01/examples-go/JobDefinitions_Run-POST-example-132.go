package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Run-POST-example-132.json
func ExampleJobDefinitionsClient_BeginRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobDefinitionsClient().BeginRun(ctx, "DataTransformation", "jobdeffromtestcode1", "ResourceGroupForSDKTest", "TestAzureSDKOperations", armhybriddatamanager.RunParameters{
		CustomerSecrets: []*armhybriddatamanager.CustomerSecret{},
		DataServiceInput: map[string]any{
			"AzureStorageType": "Blob",
			"BackupChoice":     "UseExistingLatest",
			"ContainerName":    "containerfromtest",
			"DeviceName":       "8600-SHG0997877L71FC",
			"FileNameFilter":   "*",
			"IsDirectoryMode":  false,
			"RootDirectories": []any{
				"\\",
			},
			"VolumeNames": []any{
				"TestAutomation",
			},
		},
		UserConfirmation: to.Ptr(armhybriddatamanager.UserConfirmationNotRequired),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
