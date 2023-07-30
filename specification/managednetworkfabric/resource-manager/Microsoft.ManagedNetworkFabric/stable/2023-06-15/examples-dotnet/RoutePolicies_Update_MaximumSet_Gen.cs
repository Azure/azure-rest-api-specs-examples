using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_Update_MaximumSet_Gen.json
// this example is just showing the usage of "RoutePolicies_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricRoutePolicyResource created on azure
// for more information of creating NetworkFabricRoutePolicyResource, please refer to the document of NetworkFabricRoutePolicyResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string routePolicyName = "example-routePolicy";
ResourceIdentifier networkFabricRoutePolicyResourceId = NetworkFabricRoutePolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, routePolicyName);
NetworkFabricRoutePolicyResource networkFabricRoutePolicy = client.GetNetworkFabricRoutePolicyResource(networkFabricRoutePolicyResourceId);

// invoke the operation
NetworkFabricRoutePolicyPatch patch = new NetworkFabricRoutePolicyPatch()
{
    Statements =
    {
    new RoutePolicyStatementProperties(7,new StatementConditionProperties()
    {
    RoutePolicyConditionType = RoutePolicyConditionType.Or,
    IPPrefixId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
    IPExtendedCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")
    },
    IPCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")
    },
    },new StatementActionProperties(RoutePolicyActionType.Permit)
    {
    LocalPreference = 20,
    IPCommunityProperties = new ActionIPCommunityProperties()
    {
    DeleteIPCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")
    },
    SetIPCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")
    },
    AddIPCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")
    },
    },
    IPExtendedCommunityProperties = new ActionIPExtendedCommunityProperties()
    {
    DeleteIPExtendedCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")
    },
    SetIPExtendedCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")
    },
    AddIPExtendedCommunityIds =
    {
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")
    },
    },
    })
    {
    Annotation = "annotation",
    }
    },
    Tags =
    {
    ["keyID"] = "keyValue",
    },
};
ArmOperation<NetworkFabricRoutePolicyResource> lro = await networkFabricRoutePolicy.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricRoutePolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricRoutePolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
