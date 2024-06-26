package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Operations_List-GET-example-11.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AvailableProviderOperations = armhybriddatamanager.AvailableProviderOperations{
		// 	Value: []*armhybriddatamanager.AvailableProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/write"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update Job definitions"),
		// 				Operation: to.Ptr("Creates or updates Job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/delete"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Delete Job definitions"),
		// 				Operation: to.Ptr("Delete Job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Job definitions"),
		// 				Operation: to.Ptr("Get Job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/jobDefinitions/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Job definitions"),
		// 				Operation: to.Ptr("Get Job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/listResults/action"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Get job defintions"),
		// 				Operation: to.Ptr("Get job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/run/action"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Run job defintions"),
		// 				Operation: to.Ptr("Run job definitions"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Job definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataStores/operationResults/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Operation results"),
		// 				Operation: to.Ptr("Get Operation results"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Operation results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/operationResults/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Operation results"),
		// 				Operation: to.Ptr("Get Operation results"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Operation results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/jobs/operationResults/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Operation results"),
		// 				Operation: to.Ptr("Get Operation results"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Operation results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/operationResults/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Operation results"),
		// 				Operation: to.Ptr("Get Operation results"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Operation results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/publicKeys/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Public keys"),
		// 				Operation: to.Ptr("Get Public keys"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Public keys"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Data services"),
		// 				Operation: to.Ptr("Get Data services"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Data services"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataStores/write"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update Datastores"),
		// 				Operation: to.Ptr("Creates or updates Datastores"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Datastores"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataStores/delete"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Delete Datastores"),
		// 				Operation: to.Ptr("Delete Datastores"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Datastores"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataStores/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Datastores"),
		// 				Operation: to.Ptr("Get Datastores"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Datastores"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataStoreTypes/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Datastore types"),
		// 				Operation: to.Ptr("Get Datastore types"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Datastore types"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/delete"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Delete Data managers"),
		// 				Operation: to.Ptr("Delete Data managers"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Data managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Data managers"),
		// 				Operation: to.Ptr("Get Data managers"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Data managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/write"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update Data managers"),
		// 				Operation: to.Ptr("Creates or updates Data managers"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Data managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/jobs/cancel/action"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Cancel jobs"),
		// 				Operation: to.Ptr("Cancel jobs"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/jobs/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Jobs"),
		// 				Operation: to.Ptr("Get Jobs"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/jobs/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Jobs"),
		// 				Operation: to.Ptr("Get Jobs"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobs/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read Jobs"),
		// 				Operation: to.Ptr("Get Jobs"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/dataManagers/dataServices/jobDefinitions/jobs/resume/action"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Resume jobs"),
		// 				Operation: to.Ptr("Resumes jobs"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridData/read"),
		// 			Display: &armhybriddatamanager.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Read ArmApiRes_Microsoft.HybridData"),
		// 				Operation: to.Ptr("Get ArmApiRes_Microsoft.HybridData"),
		// 				Provider: to.Ptr("Microsoft.HybridData"),
		// 				Resource: to.Ptr("ArmApiRes_Microsoft.HybridData"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 			Properties: map[string]any{
		// 			},
		// 	}},
		// }
	}
}
