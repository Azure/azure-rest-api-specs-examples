using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listTestJobStreamsByJob.json
// this example is just showing the usage of "TestJobStreams_ListByTestJob" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationRunbookResource created on azure
// for more information of creating AutomationRunbookResource, please refer to the document of AutomationRunbookResource
string subscriptionId = "51766542-3ed7-4a72-a187-0c8ab644ddab";
string resourceGroupName = "mygroup";
string automationAccountName = "ContoseAutomationAccount";
string runbookName = "Get-AzureVMTutorial";
ResourceIdentifier automationRunbookResourceId = AutomationRunbookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, runbookName);
AutomationRunbookResource automationRunbook = client.GetAutomationRunbookResource(automationRunbookResourceId);

// invoke the operation and iterate over the result
await foreach (AutomationJobStream item in automationRunbook.GetTestJobStreamsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
