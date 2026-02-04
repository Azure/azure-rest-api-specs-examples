using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayCreate.json
// this example is just showing the usage of "ServiceGateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ServiceGatewayResource
ServiceGatewayCollection collection = resourceGroupResource.GetServiceGateways();

// invoke the operation
string serviceGatewayName = "sg";
ServiceGatewayData data = new ServiceGatewayData(new AzureLocation("eastus"))
{
    VirtualNetwork = new VirtualNetworkData
    {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet"),
    },
    RouteTargetAddress = new RouteTargetAddressPropertiesFormat
    {
        Subnet = new SubnetData
        {
            Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet"),
        },
        PrivateIPAddress = "10.0.1.4",
        PrivateIPAllocationMethod = NetworkIPAllocationMethod.Static,
    },
};
ArmOperation<ServiceGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceGatewayName, data);
ServiceGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
