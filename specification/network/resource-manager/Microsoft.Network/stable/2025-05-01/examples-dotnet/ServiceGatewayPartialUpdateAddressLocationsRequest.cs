using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayPartialUpdateAddressLocationsRequest.json
// this example is just showing the usage of "ServiceGateways_UpdateAddressLocations" operation, for the dependent resources, they will have to be created separately.

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
ServiceGatewayUpdateAddressLocationsContent content = new ServiceGatewayUpdateAddressLocationsContent
{
    Action = UpdateAction.PartialUpdate,
    AddressLocations = {new ServiceGatewayAddressLocation
    {
    AddressLocation = "192.0.0.1",
    AddressUpdateAction = AddressUpdateAction.FullUpdate,
    Addresses = {new ServiceGatewayAddress
    {
    Address = "10.0.0.4",
    Services = {"Service1"},
    }},
    }, new ServiceGatewayAddressLocation
    {
    AddressLocation = "192.0.0.2",
    AddressUpdateAction = AddressUpdateAction.PartialUpdate,
    Addresses = {new ServiceGatewayAddress
    {
    Address = "10.0.0.5",
    Services = {"Service2"},
    }, new ServiceGatewayAddress
    {
    Address = "10.0.0.6",
    }},
    }, new ServiceGatewayAddressLocation
    {
    AddressLocation = "192.0.0.3",
    }},
};
await serviceGateway.UpdateAddressLocationsAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
