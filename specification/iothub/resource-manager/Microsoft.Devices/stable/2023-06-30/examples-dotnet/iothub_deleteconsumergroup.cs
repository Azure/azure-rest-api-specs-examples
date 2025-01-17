using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.IotHub;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_deleteconsumergroup.json
// this example is just showing the usage of "IotHubResource_DeleteEventHubConsumerGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubConsumerGroupInfoResource created on azure
// for more information of creating EventHubConsumerGroupInfoResource, please refer to the document of EventHubConsumerGroupInfoResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string resourceName = "testHub";
string eventHubEndpointName = "events";
string name = "test";
ResourceIdentifier eventHubConsumerGroupInfoResourceId = EventHubConsumerGroupInfoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, eventHubEndpointName, name);
EventHubConsumerGroupInfoResource eventHubConsumerGroupInfo = client.GetEventHubConsumerGroupInfoResource(eventHubConsumerGroupInfoResourceId);

// invoke the operation
await eventHubConsumerGroupInfo.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
