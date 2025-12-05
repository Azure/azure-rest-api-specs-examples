using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/NameSpaces/PrivateEndPointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusPrivateEndpointConnectionResource created on azure
// for more information of creating ServiceBusPrivateEndpointConnectionResource, please refer to the document of ServiceBusPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-3285";
string privateEndpointConnectionName = "928c44d5-b7c6-423b-b6fa-811e0c27b3e0";
ResourceIdentifier serviceBusPrivateEndpointConnectionResourceId = ServiceBusPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
ServiceBusPrivateEndpointConnectionResource serviceBusPrivateEndpointConnection = client.GetServiceBusPrivateEndpointConnectionResource(serviceBusPrivateEndpointConnectionResourceId);

// invoke the operation
await serviceBusPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
