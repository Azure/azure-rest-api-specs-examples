package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/runbook/replaceRunbookDraftContent.json
func ExampleRunbookDraftClient_BeginReplaceContent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRunbookDraftClient().BeginReplaceContent(ctx, "rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", "<#\r\n        .DESCRIPTION \r\n            An example runbook which prints out the first10 Azure VMs in your subscription (ordered alphabetically).\r\n            For more information about how this runbook authenticates to your Azure subscription, see our documentation here: http: //aka.ms/fxu3mn\r\n\r\n        .NOTES\r\n            AUTHOR: Azure Automation Team\r\n            LASTEDIT: Mar27,\r\n            2015\r\n    #>\r\n    workflow Get-AzureVMTutorial{\r\n        #The name of the Automation Credential Asset this runbook will use to authenticate to Azure.\r\n        $CredentialAssetName = 'DefaultAzureCredential'\r\n\r\n        #Get the credential with the above name from the Automation Asset store\r\n        $Cred = Get-AutomationPSCredential -Name $CredentialAssetName\r\n        if(!$Cred){\r\n            Throw\"Could not find an Automation Credential Asset named '${CredentialAssetName}'. Make sure you have created one in this Automation Account.\"\r\n                }\r\n\r\n        #Connect to your Azure Account\r\n        $Account = Add-AzureAccount -Credential $Cred\r\n        if(!$Account){\r\n            Throw\"Could not authenticate to Azure using the credential asset '${CredentialAssetName}'. Make sure the user name and password are correct.\"\r\n                }\r\n\r\n        #TODO (optional): pick the right subscription to use. Without this line, the default subscription for your Azure Account will be used.\r\n        #Select-AzureSubscription -SubscriptionName\"TODO: your Azure subscription name here\"\r\n        \r\n        #Get all the VMs you have in your Azure subscription\r\n        $VMs = Get-AzureVM\r\n\r\n        #Print out up to10 of those VMs\r\n        if(!$VMs){\r\n            Write-Output\"No VMs were found in your subscription.\"\r\n                } else{\r\n            Write-Output $VMs[0..9\r\n                    ]\r\n                }\r\n            }", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
