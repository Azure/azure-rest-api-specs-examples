package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ResourceProviderOperationCollection = armblockchain.ResourceProviderOperationCollection{
		// 	Value: []*armblockchain.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/register/action"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Blockchain Resource Provider."),
		// 				Operation: to.Ptr("Register the Blockchain Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets or Lists existing Blockchain Member(s)."),
		// 				Operation: to.Ptr("Get or List Blockchain Member(s)"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Members"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/write"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Creates or Updates a Blockchain Member."),
		// 				Operation: to.Ptr("Create or Update Blockchain Member"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Members"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/delete"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Blockchain Member."),
		// 				Operation: to.Ptr("Delete Blockchain Member"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Members"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets or Lists existing Blockchain Member Transaction Node(s)."),
		// 				Operation: to.Ptr("Get or List Blockchain Member Transaction Node(s)"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Member Transaction Nodes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes/write"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Creates or Updates a Blockchain Member Transaction Node."),
		// 				Operation: to.Ptr("Create or Update Blockchain Member Transaction Node"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Member Transaction Nodes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes/delete"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Blockchain Member Transaction Node."),
		// 				Operation: to.Ptr("Delete Blockchain Member Transaction Node"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Member Transaction Nodes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes/connect/action"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Connects to a Blockchain Member Transaction Node."),
		// 				Operation: to.Ptr("Connect Blockchain Member Transaction Node"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Blockchain Member Transaction Nodes"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/locations/blockchainMemberOperationResults/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets the Operation Results of Blockchain Members."),
		// 				Operation: to.Ptr("Get Blockchain Member Operation Results"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/locations/checkNameAvailability/action"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Checks that resource name is valid and is not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/operations/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("List all Operations in Microsoft Blockchain Resource Provider."),
		// 				Operation: to.Ptr("List all Operations"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for Microsoft Blockchain"),
		// 				Operation: to.Ptr("Read Microsoft Blockchain log definitions"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("The log definition of Microsoft Blockchain"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("The display resource of diagnostic settings of Microsoft Blockchain"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Blockchain/blockchainMembers/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armblockchain.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Blockchain"),
		// 				Resource: to.Ptr("The display resource of diagnostic settings of Microsoft Blockchain"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 	}},
		// }
	}
}
