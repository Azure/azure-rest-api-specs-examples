using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/PrivateEndpointDnsZoneGroupGet.json
// this example is just showing the usage of "PrivateDnsZoneGroups_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this PrivateDnsZoneGroupResource
PrivateDnsZoneGroupCollection collection = privateEndpoint.GetPrivateDnsZoneGroups();

// invoke the operation
string privateDnsZoneGroupName = "testPdnsgroup";
bool result = await collection.ExistsAsync(privateDnsZoneGroupName);

Console.WriteLine($"Succeeded: {result}");
