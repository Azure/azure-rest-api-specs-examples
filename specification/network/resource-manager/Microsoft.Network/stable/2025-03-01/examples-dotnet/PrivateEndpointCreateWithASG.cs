using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateEndpointCreateWithASG.json
// this example is just showing the usage of "PrivateEndpoints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateEndpointResource created on azure
// for more information of creating PrivateEndpointResource, please refer to the document of PrivateEndpointResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string privateEndpointName = "testPe";
ResourceIdentifier privateEndpointResourceId = PrivateEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateEndpointName);
PrivateEndpointResource privateEndpoint = client.GetPrivateEndpointResource(privateEndpointResourceId);

// invoke the operation
PrivateEndpointData data = new PrivateEndpointData
{
    Subnet = new SubnetData
    {
        Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
    },
    PrivateLinkServiceConnections = {new NetworkPrivateLinkServiceConnection
    {
    PrivateLinkServiceId = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
    GroupIds = {"groupIdFromResource"},
    RequestMessage = "Please approve my connection.",
    }},
    ApplicationSecurityGroups = {new ApplicationSecurityGroupData
    {
    Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1"),
    }},
    Location = new AzureLocation("eastus2euap"),
};
ArmOperation<PrivateEndpointResource> lro = await privateEndpoint.UpdateAsync(WaitUntil.Completed, data);
PrivateEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
