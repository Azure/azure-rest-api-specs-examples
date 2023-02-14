using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getModule.json
// this example is just showing the usage of "Module_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
AutomationAccountModuleResource result = await automationAccountModule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationModuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
