using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ManagedNetwork;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkGroup/ManagedNetworkGroupsPut.json
// this example is just showing the usage of "ManagedNetworkGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedNetworkGroupResource created on azure
// for more information of creating ManagedNetworkGroupResource, please refer to the document of ManagedNetworkGroupResource
string subscriptionId = "subscriptionA";
string resourceGroupName = "myResourceGroup";
string managedNetworkName = "myManagedNetwork";
string managedNetworkGroupName = "myManagedNetworkGroup1";
ResourceIdentifier managedNetworkGroupResourceId = ManagedNetworkGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedNetworkName, managedNetworkGroupName);
ManagedNetworkGroupResource managedNetworkGroup = client.GetManagedNetworkGroupResource(managedNetworkGroupResourceId);

// invoke the operation
ManagedNetworkGroupData data = new ManagedNetworkGroupData
{
    ManagementGroups = { },
    Subscriptions = { },
    VirtualNetworks = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
    }, new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
    }},
    Subnets = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA/subnets/subnetA"),
    }},
};
ArmOperation<ManagedNetworkGroupResource> lro = await managedNetworkGroup.UpdateAsync(WaitUntil.Completed, data);
ManagedNetworkGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedNetworkGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
