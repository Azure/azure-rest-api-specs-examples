using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L3IsolationDomains_Create_MaximumSet_Gen.json
// this example is just showing the usage of "L3IsolationDomains_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFabricL3IsolationDomainResource
NetworkFabricL3IsolationDomainCollection collection = resourceGroupResource.GetNetworkFabricL3IsolationDomains();

// invoke the operation
string l3IsolationDomainName = "example-l3domain";
NetworkFabricL3IsolationDomainData data = new NetworkFabricL3IsolationDomainData(new AzureLocation("eastus"), new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"))
{
    Annotation = "annotation",
    RedistributeConnectedSubnets = RedistributeConnectedSubnet.True,
    RedistributeStaticRoutes = RedistributeStaticRoute.False,
    AggregateRouteConfiguration = new AggregateRouteConfiguration
    {
        IPv4Routes = { new AggregateRoute("10.0.0.0/24") },
        IPv6Routes = { new AggregateRoute("3FFE:FFFF:0:CD30::a0/29") },
    },
    ConnectedSubnetRoutePolicy = new ConnectedSubnetRoutePolicy
    {
        ExportRoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
        ExportRoutePolicy = new L3ExportRoutePolicy
        {
            ExportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
            ExportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
        },
    },
    Tags =
    {
    ["keyID"] = "KeyValue"
    },
};
ArmOperation<NetworkFabricL3IsolationDomainResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, l3IsolationDomainName, data);
NetworkFabricL3IsolationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricL3IsolationDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
