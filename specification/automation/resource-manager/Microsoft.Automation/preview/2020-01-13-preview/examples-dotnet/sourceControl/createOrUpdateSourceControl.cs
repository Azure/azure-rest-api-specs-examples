using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/createOrUpdateSourceControl.json
// this example is just showing the usage of "SourceControl_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "sampleAccount9";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// get the collection of this AutomationSourceControlResource
AutomationSourceControlCollection collection = automationAccount.GetAutomationSourceControls();

// invoke the operation
string sourceControlName = "sampleSourceControl";
AutomationSourceControlCreateOrUpdateContent content = new AutomationSourceControlCreateOrUpdateContent()
{
    RepoUri = new Uri("https://sampleUser.visualstudio.com/myProject/_git/myRepository"),
    Branch = "master",
    FolderPath = "/folderOne/folderTwo",
    IsAutoSyncEnabled = true,
    IsAutoPublishRunbookEnabled = true,
    SourceType = SourceControlSourceType.VsoGit,
    SecurityToken = new SourceControlSecurityTokenProperties()
    {
        AccessToken = "3a326f7a0dcd343ea58fee21f2fd5fb4c1234567",
        TokenType = SourceControlTokenType.PersonalAccessToken,
    },
    Description = "my description",
};
ArmOperation<AutomationSourceControlResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sourceControlName, content);
AutomationSourceControlResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationSourceControlData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
