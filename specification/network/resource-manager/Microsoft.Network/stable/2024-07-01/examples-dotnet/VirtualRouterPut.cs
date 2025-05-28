using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualRouterPut.json
// this example is just showing the usage of "VirtualRouters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualRouterResource created on azure
// for more information of creating VirtualRouterResource, please refer to the document of VirtualRouterResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualRouterName = "virtualRouter";
ResourceIdentifier virtualRouterResourceId = VirtualRouterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualRouterName);
VirtualRouterResource virtualRouter = client.GetVirtualRouterResource(virtualRouterResourceId);

// invoke the operation
VirtualRouterData data = new VirtualRouterData
{
    HostedGatewayId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway"),
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<VirtualRouterResource> lro = await virtualRouter.UpdateAsync(WaitUntil.Completed, data);
VirtualRouterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualRouterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
