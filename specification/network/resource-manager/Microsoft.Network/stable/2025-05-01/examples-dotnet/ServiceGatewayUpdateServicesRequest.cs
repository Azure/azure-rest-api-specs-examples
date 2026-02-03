using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayUpdateServicesRequest.json
// this example is just showing the usage of "ServiceGateways_UpdateServices" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceGatewayResource created on azure
// for more information of creating ServiceGatewayResource, please refer to the document of ServiceGatewayResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceGatewayName = "sg";
ResourceIdentifier serviceGatewayResourceId = ServiceGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceGatewayName);
ServiceGatewayResource serviceGateway = client.GetServiceGatewayResource(serviceGatewayResourceId);

// invoke the operation
ServiceGatewayUpdateServicesContent content = new ServiceGatewayUpdateServicesContent
{
    Action = ServiceUpdateAction.FullUpdate,
    ServiceRequests = {new ServiceGatewayServiceRequest
    {
    Service = new ServiceGatewayService
    {
    Name = "Service1",
    ServiceType = ServiceType.Inbound,
    IsDefault = true,
    LoadBalancerBackendPools = {new BackendAddressPoolData
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/be1"),
    }},
    PublicNatGatewayId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-natGateway",
    },
    }, new ServiceGatewayServiceRequest
    {
    IsDelete = true,
    Service = new ServiceGatewayService
    {
    Name = "Service2",
    ServiceType = ServiceType.Outbound,
    IsDefault = false,
    },
    }},
};
await serviceGateway.UpdateServicesAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
