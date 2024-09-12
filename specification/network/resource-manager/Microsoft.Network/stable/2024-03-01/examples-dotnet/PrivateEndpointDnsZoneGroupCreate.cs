using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/PrivateEndpointDnsZoneGroupCreate.json
// this example is just showing the usage of "PrivateDnsZoneGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateDnsZoneGroupResource created on azure
// for more information of creating PrivateDnsZoneGroupResource, please refer to the document of PrivateDnsZoneGroupResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string privateEndpointName = "testPe";
string privateDnsZoneGroupName = "testPdnsgroup";
ResourceIdentifier privateDnsZoneGroupResourceId = PrivateDnsZoneGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateEndpointName, privateDnsZoneGroupName);
PrivateDnsZoneGroupResource privateDnsZoneGroup = client.GetPrivateDnsZoneGroupResource(privateDnsZoneGroupResourceId);

// invoke the operation
PrivateDnsZoneGroupData data = new PrivateDnsZoneGroupData()
{
    PrivateDnsZoneConfigs =
    {
    new PrivateDnsZoneConfig()
    {
    PrivateDnsZoneId = "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone1.com",
    }
    },
};
ArmOperation<PrivateDnsZoneGroupResource> lro = await privateDnsZoneGroup.UpdateAsync(WaitUntil.Completed, data);
PrivateDnsZoneGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateDnsZoneGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
