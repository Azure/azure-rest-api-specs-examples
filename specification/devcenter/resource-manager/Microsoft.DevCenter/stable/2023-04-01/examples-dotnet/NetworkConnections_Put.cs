using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_Put.json
// this example is just showing the usage of "NetworkConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DevCenterNetworkConnectionResource
DevCenterNetworkConnectionCollection collection = resourceGroupResource.GetDevCenterNetworkConnections();

// invoke the operation
string networkConnectionName = "uswest3network";
DevCenterNetworkConnectionData data = new DevCenterNetworkConnectionData(new AzureLocation("centralus"))
{
    SubnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ExampleRG/providers/Microsoft.Network/virtualNetworks/ExampleVNet/subnets/default"),
    DomainName = "mydomaincontroller.local",
    DomainUsername = "testuser@mydomaincontroller.local",
    DomainPassword = "Password value for user",
    NetworkingResourceGroupName = "NetworkInterfaces",
    DomainJoinType = DomainJoinType.HybridAadJoin,
};
ArmOperation<DevCenterNetworkConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkConnectionName, data);
DevCenterNetworkConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevCenterNetworkConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
