using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotCentral;
using Azure.ResourceManager.IotCentral.Models;

// Generated from example definition: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotCentralPrivateEndpointConnectionResource created on azure
// for more information of creating IotCentralPrivateEndpointConnectionResource, please refer to the document of IotCentralPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string resourceName = "myIoTCentralApp";
string privateEndpointConnectionName = "myIoTCentralAppEndpoint";
ResourceIdentifier iotCentralPrivateEndpointConnectionResourceId = IotCentralPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
IotCentralPrivateEndpointConnectionResource iotCentralPrivateEndpointConnection = client.GetIotCentralPrivateEndpointConnectionResource(iotCentralPrivateEndpointConnectionResourceId);

// invoke the operation
await iotCentralPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
