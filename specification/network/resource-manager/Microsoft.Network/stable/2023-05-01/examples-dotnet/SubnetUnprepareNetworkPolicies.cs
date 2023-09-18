using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/SubnetUnprepareNetworkPolicies.json
// this example is just showing the usage of "Subnets_UnprepareNetworkPolicies" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubnetResource created on azure
// for more information of creating SubnetResource, please refer to the document of SubnetResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkName = "test-vnet";
string subnetName = "subnet1";
ResourceIdentifier subnetResourceId = SubnetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName, subnetName);
SubnetResource subnet = client.GetSubnetResource(subnetResourceId);

// invoke the operation
UnprepareNetworkPoliciesContent content = new UnprepareNetworkPoliciesContent()
{
    ServiceName = "Microsoft.Sql/managedInstances",
};
await subnet.UnprepareNetworkPoliciesAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
