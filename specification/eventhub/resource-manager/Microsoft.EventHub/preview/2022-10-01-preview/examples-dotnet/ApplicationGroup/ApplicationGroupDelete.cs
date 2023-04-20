using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/ApplicationGroup/ApplicationGroupDelete.json
// this example is just showing the usage of "ApplicationGroup_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsApplicationGroupResource created on azure
// for more information of creating EventHubsApplicationGroupResource, please refer to the document of EventHubsApplicationGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosotest";
string namespaceName = "contoso-ua-test-eh-system-1";
string applicationGroupName = "appGroup1";
ResourceIdentifier eventHubsApplicationGroupResourceId = EventHubsApplicationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, applicationGroupName);
EventHubsApplicationGroupResource eventHubsApplicationGroup = client.GetEventHubsApplicationGroupResource(eventHubsApplicationGroupResourceId);

// invoke the operation
await eventHubsApplicationGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
