package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_GenerateArmTemplate.json
func ExampleArtifactsClient_GenerateArmTemplate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactsClient().GenerateArmTemplate(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", "{artifactName}", armdevtestlabs.GenerateArmTemplateRequest{
		FileUploadOptions:  to.Ptr(armdevtestlabs.FileUploadOptionsNone),
		Location:           to.Ptr("{location}"),
		VirtualMachineName: to.Ptr("{vmName}"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmTemplateInfo = armdevtestlabs.ArmTemplateInfo{
	// 	Parameters: map[string]any{
	// 		"extensionName":map[string]any{
	// 			"value": "{vmName}/CustomScriptExtension",
	// 		},
	// 		"location":map[string]any{
	// 			"value": "{location}",
	// 		},
	// 	},
	// 	Template: map[string]any{
	// 		"$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json",
	// 		"contentVersion": "1.0.0.0",
	// 		"parameters":map[string]any{
	// 			"extensionName":map[string]any{
	// 				"type": "string",
	// 			},
	// 			"location":map[string]any{
	// 				"type": "string",
	// 			},
	// 		},
	// 		"resources":[]any{
	// 			map[string]any{
	// 				"name": "[parameters('extensionName')]",
	// 				"type": "Microsoft.Compute/virtualMachines/extensions",
	// 				"apiVersion": "2015-06-15",
	// 				"location": "[parameters('location')]",
	// 				"properties":map[string]any{
	// 					"type": "CustomScriptExtension",
	// 					"autoUpgradeMinorVersion": "true",
	// 					"forceUpdateTag": "15/10/2018 00:00:00 +00:00",
	// 					"protectedSettings":map[string]any{
	// 						"commandToExecute": "[concat('cd {MsDtlScriptFolder}/scripts && ', variables('_commandToExecute'))]",
	// 					},
	// 					"publisher": "Microsoft.Compute",
	// 					"settings":map[string]any{
	// 						"commandToExecute": "",
	// 						"fileUris":[]any{
	// 							"{MsDtlArtifactFileUris}",
	// 						},
	// 					},
	// 					"typeHandlerVersion": "1.9",
	// 				},
	// 			},
	// 		},
	// 		"variables":map[string]any{
	// 			"_commandToExecute": "{commandToExecute}.",
	// 		},
	// 	},
	// }
}
