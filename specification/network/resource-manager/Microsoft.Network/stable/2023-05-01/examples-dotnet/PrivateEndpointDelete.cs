using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateEndpointDelete.json
// this example is just showing the usage of "PrivateEndpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateEndpointResource created on azure
// for more information of creating PrivateEndpointResource, please refer to the document of PrivateEndpointResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string privateEndpointName = "testPe";
ResourceIdentifier privateEndpointResourceId = PrivateEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateEndpointName);
PrivateEndpointResource privateEndpoint = client.GetPrivateEndpointResource(privateEndpointResourceId);

// invoke the operation
await privateEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
