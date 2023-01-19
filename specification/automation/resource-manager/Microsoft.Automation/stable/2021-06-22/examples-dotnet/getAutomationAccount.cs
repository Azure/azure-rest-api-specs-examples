using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getAutomationAccount.json
// this example is just showing the usage of "AutomationAccount_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AutomationAccountResource
AutomationAccountCollection collection = resourceGroupResource.GetAutomationAccounts();

// invoke the operation
string automationAccountName = "myAutomationAccount9";
bool result = await collection.ExistsAsync(automationAccountName);

Console.WriteLine($"Succeeded: {result}");
