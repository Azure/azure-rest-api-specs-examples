using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/deleteRunbook.json
// this example is just showing the usage of "Runbook_Delete" operation, for the dependent resources, they will have to be created separately.

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
await automationRunbook.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
