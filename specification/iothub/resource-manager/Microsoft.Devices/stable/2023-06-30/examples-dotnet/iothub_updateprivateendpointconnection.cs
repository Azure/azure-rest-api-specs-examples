using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.IotHub;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_updateprivateendpointconnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotHubPrivateEndpointConnectionResource created on azure
// for more information of creating IotHubPrivateEndpointConnectionResource, please refer to the document of IotHubPrivateEndpointConnectionResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string resourceName = "testHub";
string privateEndpointConnectionName = "myPrivateEndpointConnection";
ResourceIdentifier iotHubPrivateEndpointConnectionResourceId = IotHubPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
IotHubPrivateEndpointConnectionResource iotHubPrivateEndpointConnection = client.GetIotHubPrivateEndpointConnectionResource(iotHubPrivateEndpointConnectionResourceId);

// invoke the operation
IotHubPrivateEndpointConnectionData data = new IotHubPrivateEndpointConnectionData(new IotHubPrivateEndpointConnectionProperties(new IotHubPrivateLinkServiceConnectionState(IotHubPrivateLinkServiceConnectionStatus.Approved, "Approved by johndoe@contoso.com")));
ArmOperation<IotHubPrivateEndpointConnectionResource> lro = await iotHubPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
IotHubPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotHubPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
