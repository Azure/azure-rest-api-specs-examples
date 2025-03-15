using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listActivitiesByModule.json
// this example is just showing the usage of "Activity_ListByModule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountModuleResource created on azure
// for more information of creating AutomationAccountModuleResource, please refer to the document of AutomationAccountModuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string moduleName = "OmsCompositeResources";
ResourceIdentifier automationAccountModuleResourceId = AutomationAccountModuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, moduleName);
AutomationAccountModuleResource automationAccountModule = client.GetAutomationAccountModuleResource(automationAccountModuleResourceId);

// invoke the operation and iterate over the result
await foreach (AutomationActivity item in automationAccountModule.GetActivitiesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
