using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
// this example is just showing the usage of "Runbook_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationRunbookResource created on azure
// for more information of creating AutomationRunbookResource, please refer to the document of AutomationRunbookResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "ContoseAutomationAccount";
string runbookName = "Get-AzureVMTutorial";
ResourceIdentifier automationRunbookResourceId = AutomationRunbookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, runbookName);
AutomationRunbookResource automationRunbook = client.GetAutomationRunbookResource(automationRunbookResourceId);

// invoke the operation
AutomationRunbookPatch patch = new AutomationRunbookPatch
{
    Description = "Updated Description of the Runbook",
    IsLogVerboseEnabled = false,
    IsLogProgressEnabled = true,
    LogActivityTrace = 1,
};
AutomationRunbookResource result = await automationRunbook.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationRunbookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
