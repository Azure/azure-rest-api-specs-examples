using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createRunbookAsDraft.json
// this example is just showing the usage of "Runbook_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "ContoseAutomationAccount";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// get the collection of this AutomationRunbookResource
AutomationRunbookCollection collection = automationAccount.GetAutomationRunbooks();

// invoke the operation
string runbookName = "Get-AzureVMTutorial";
AutomationRunbookCreateOrUpdateContent content = new AutomationRunbookCreateOrUpdateContent(AutomationRunbookType.PowerShellWorkflow)
{
    Name = "Get-AzureVMTutorial",
    Location = new AzureLocation("East US 2"),
    Tags =
    {
    ["tag01"] = "value01",
    ["tag02"] = "value02"
    },
    IsLogVerboseEnabled = false,
    IsLogProgressEnabled = false,
    Draft = new AutomationRunbookDraft(),
    Description = "Description of the Runbook",
};
ArmOperation<AutomationRunbookResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, runbookName, content);
AutomationRunbookResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationRunbookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
