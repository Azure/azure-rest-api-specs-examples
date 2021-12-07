Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblueprint%2Farmblueprint%2Fv0.1.0/sdk/resourcemanager/blueprint/armblueprint/README.md) on how to add the SDK to your project and authenticate.

```go
package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// x-ms-original-file: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Create_SystemAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblueprint.NewAssignmentsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-scope>",
		"<assignment-name>",
		armblueprint.Assignment{
			TrackedResource: armblueprint.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armblueprint.ManagedServiceIdentity{
				Type: armblueprint.ManagedServiceIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armblueprint.AssignmentProperties{
				BlueprintResourcePropertiesBase: armblueprint.BlueprintResourcePropertiesBase{
					Description: to.StringPtr("<description>"),
				},
				BlueprintID: to.StringPtr("<blueprint-id>"),
				Parameters: map[string]*armblueprint.ParameterValue{
					"costCenter": {
						Value: map[string]interface{}{
							"0":  "C",
							"1":  "o",
							"2":  "n",
							"3":  "t",
							"4":  "o",
							"5":  "s",
							"6":  "o",
							"7":  "/",
							"8":  "O",
							"9":  "n",
							"10": "l",
							"11": "i",
							"12": "n",
							"13": "e",
							"14": "/",
							"15": "S",
							"16": "h",
							"17": "o",
							"18": "p",
							"19": "p",
							"20": "i",
							"21": "n",
							"22": "g",
							"23": "/",
							"24": "P",
							"25": "r",
							"26": "o",
							"27": "d",
							"28": "u",
							"29": "c",
							"30": "t",
							"31": "i",
							"32": "o",
							"33": "n",
						},
					},
					"owners": {
						Value: map[string]interface{}{
							"0": "johnDoe@contoso.com",
							"1": "johnsteam@contoso.com",
						},
					},
					"storageAccountType": {
						Value: map[string]interface{}{
							"0":  "S",
							"1":  "t",
							"2":  "a",
							"3":  "n",
							"4":  "d",
							"5":  "a",
							"6":  "r",
							"7":  "d",
							"8":  "_",
							"9":  "L",
							"10": "R",
							"11": "S",
						},
					},
				},
				ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
					"storageRG": {
						Name:     to.StringPtr("<name>"),
						Location: to.StringPtr("<location>"),
					},
				},
				Scope: to.StringPtr("<scope>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Assignment.ID: %s\n", *res.ID)
}
```
