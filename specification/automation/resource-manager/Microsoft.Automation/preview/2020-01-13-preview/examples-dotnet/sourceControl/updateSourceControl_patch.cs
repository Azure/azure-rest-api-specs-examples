using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/updateSourceControl_patch.json
// this example is just showing the usage of "SourceControl_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationSourceControlResource created on azure
// for more information of creating AutomationSourceControlResource, please refer to the document of AutomationSourceControlResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "sampleAccount9";
string sourceControlName = "sampleSourceControl";
ResourceIdentifier automationSourceControlResourceId = AutomationSourceControlResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, sourceControlName);
AutomationSourceControlResource automationSourceControl = client.GetAutomationSourceControlResource(automationSourceControlResourceId);

// invoke the operation
AutomationSourceControlPatch patch = new AutomationSourceControlPatch
{
    Branch = "master",
    FolderPath = "/folderOne/folderTwo",
    IsAutoSyncEnabled = true,
    IsAutoPublishRunbookEnabled = true,
    SecurityToken = new SourceControlSecurityTokenProperties
    {
        AccessToken = "3a326f7a0dcd343ea58fee21f2fd5fb4c1234567",
        TokenType = SourceControlTokenType.PersonalAccessToken,
    },
    Description = "my description",
};
AutomationSourceControlResource result = await automationSourceControl.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationSourceControlData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
