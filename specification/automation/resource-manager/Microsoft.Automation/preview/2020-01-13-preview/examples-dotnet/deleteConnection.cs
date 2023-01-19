using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteConnection.json
// this example is just showing the usage of "Connection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationConnectionResource created on azure
// for more information of creating AutomationConnectionResource, please refer to the document of AutomationConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount28";
string connectionName = "myConnection";
ResourceIdentifier automationConnectionResourceId = AutomationConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, connectionName);
AutomationConnectionResource automationConnection = client.GetAutomationConnectionResource(automationConnectionResourceId);

// invoke the operation
await automationConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
