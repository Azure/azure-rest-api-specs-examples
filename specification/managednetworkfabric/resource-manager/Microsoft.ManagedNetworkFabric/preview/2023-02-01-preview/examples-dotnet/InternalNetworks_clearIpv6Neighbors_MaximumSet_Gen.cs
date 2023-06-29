using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_clearIpv6Neighbors_MaximumSet_Gen.json
// this example is just showing the usage of "InternalNetworks_clearIpv6Neighbors" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InternalNetworkResource created on azure
// for more information of creating InternalNetworkResource, please refer to the document of InternalNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "example-l3domain";
string internalNetworkName = "example-internalnetwork";
ResourceIdentifier internalNetworkResourceId = InternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, internalNetworkName);
InternalNetworkResource internalNetwork = client.GetInternalNetworkResource(internalNetworkResourceId);

// invoke the operation
EnableDisableOnResources body = new EnableDisableOnResources()
{
    ResourceIds =
    {
    "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"
    },
};
await internalNetwork.ClearIPv6NeighborsAsync(WaitUntil.Completed, body);

Console.WriteLine($"Succeeded");
