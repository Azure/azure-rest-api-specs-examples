using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricControllers_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabricControllers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricControllerResource created on azure
// for more information of creating NetworkFabricControllerResource, please refer to the document of NetworkFabricControllerResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkFabricControllerName = "example-networkController";
ResourceIdentifier networkFabricControllerResourceId = NetworkFabricControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricControllerName);
NetworkFabricControllerResource networkFabricController = client.GetNetworkFabricControllerResource(networkFabricControllerResourceId);

// invoke the operation
NetworkFabricControllerPatch patch = new NetworkFabricControllerPatch()
{
    InfrastructureExpressRouteConnections =
    {
    new ExpressRouteConnectionInformation(new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"))
    {
    ExpressRouteAuthorizationKey = "xxxxxxx",
    }
    },
    WorkloadExpressRouteConnections =
    {
    new ExpressRouteConnectionInformation(new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"))
    {
    ExpressRouteAuthorizationKey = "xxxxxxx",
    }
    },
};
ArmOperation<NetworkFabricControllerResource> lro = await networkFabricController.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricControllerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricControllerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
