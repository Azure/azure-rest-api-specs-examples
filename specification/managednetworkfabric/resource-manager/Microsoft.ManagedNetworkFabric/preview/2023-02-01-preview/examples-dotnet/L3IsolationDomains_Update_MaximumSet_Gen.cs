using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_Update_MaximumSet_Gen.json
// this example is just showing the usage of "L3IsolationDomains_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this L3IsolationDomainResource created on azure
// for more information of creating L3IsolationDomainResource, please refer to the document of L3IsolationDomainResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "example-l3domain";
ResourceIdentifier l3IsolationDomainResourceId = L3IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName);
L3IsolationDomainResource l3IsolationDomain = client.GetL3IsolationDomainResource(l3IsolationDomainResourceId);

// invoke the operation
L3IsolationDomainPatch patch = new L3IsolationDomainPatch()
{
    RedistributeConnectedSubnets = RedistributeConnectedSubnet.True,
    RedistributeStaticRoutes = RedistributeStaticRoute.False,
    AggregateRouteConfiguration = new AggregateRouteConfiguration()
    {
        IPv4Routes =
        {
        new AggregateRoute()
        {
        Prefix = "10.0.0.0/24",
        }
        },
        IPv6Routes =
        {
        new AggregateRoute()
        {
        Prefix = "3FFE:FFFF:0:CD30::a0/29",
        }
        },
    },
    Description = "creating L3 isolation domain",
    ConnectedSubnetRoutePolicy = new L3IsolationDomainPatchPropertiesConnectedSubnetRoutePolicy()
    {
        ExportRoutePolicyId = "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
    },
};
ArmOperation<L3IsolationDomainResource> lro = await l3IsolationDomain.UpdateAsync(WaitUntil.Completed, patch);
L3IsolationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
L3IsolationDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
