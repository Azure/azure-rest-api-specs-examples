using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VirtualHubBgpConnectionGet.json
// this example is just showing the usage of "VirtualHubBgpConnection_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BgpConnectionResource created on azure
// for more information of creating BgpConnectionResource, please refer to the document of BgpConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "hub1";
string connectionName = "conn1";
ResourceIdentifier bgpConnectionResourceId = BgpConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, connectionName);
BgpConnectionResource bgpConnection = client.GetBgpConnectionResource(bgpConnectionResourceId);

// invoke the operation
BgpConnectionResource result = await bgpConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BgpConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
