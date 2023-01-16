using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;
using Azure.ResourceManager.ServiceBus.Models;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionGet.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusPrivateEndpointConnectionResource created on azure
// for more information of creating ServiceBusPrivateEndpointConnectionResource, please refer to the document of ServiceBusPrivateEndpointConnectionResource
string subscriptionId = "subID";
string resourceGroupName = "SDK-ServiceBus-4794";
string namespaceName = "sdk-Namespace-5828";
string privateEndpointConnectionName = "privateEndpointConnectionName";
ResourceIdentifier serviceBusPrivateEndpointConnectionResourceId = ServiceBusPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
ServiceBusPrivateEndpointConnectionResource serviceBusPrivateEndpointConnection = client.GetServiceBusPrivateEndpointConnectionResource(serviceBusPrivateEndpointConnectionResourceId);

// invoke the operation
ServiceBusPrivateEndpointConnectionResource result = await serviceBusPrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
