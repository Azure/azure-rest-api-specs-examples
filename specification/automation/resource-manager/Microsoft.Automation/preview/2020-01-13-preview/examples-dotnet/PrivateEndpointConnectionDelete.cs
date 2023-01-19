using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateEndpointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationPrivateEndpointConnectionResource created on azure
// for more information of creating AutomationPrivateEndpointConnectionResource, please refer to the document of AutomationPrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "rg1";
string automationAccountName = "ddb1";
string privateEndpointConnectionName = "privateEndpointConnectionName";
ResourceIdentifier automationPrivateEndpointConnectionResourceId = AutomationPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, privateEndpointConnectionName);
AutomationPrivateEndpointConnectionResource automationPrivateEndpointConnection = client.GetAutomationPrivateEndpointConnectionResource(automationPrivateEndpointConnectionResourceId);

// invoke the operation
await automationPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
