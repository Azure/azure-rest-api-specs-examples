using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/SubnetGet.json
// this example is just showing the usage of "Subnets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubnetResource created on azure
// for more information of creating SubnetResource, please refer to the document of SubnetResource
string subscriptionId = "subid";
string resourceGroupName = "subnet-test";
string virtualNetworkName = "vnetname";
string subnetName = "subnet1";
ResourceIdentifier subnetResourceId = SubnetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName, subnetName);
SubnetResource subnet = client.GetSubnetResource(subnetResourceId);

// invoke the operation
SubnetResource result = await subnet.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SubnetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
