using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupCreate.json
// this example is just showing the usage of "SimGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this MobileNetworkSimGroupResource
MobileNetworkSimGroupCollection collection = resourceGroupResource.GetMobileNetworkSimGroups();

// invoke the operation
string simGroupName = "testSimGroup";
MobileNetworkSimGroupData data = new MobileNetworkSimGroupData(new AzureLocation("eastus"))
{
    UserAssignedIdentity = new MobileNetworkManagedServiceIdentity(MobileNetworkManagedServiceIdentityType.UserAssigned)
    {
        UserAssignedIdentities =
        {
        ["/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity"] = new UserAssignedIdentity(),
        },
    },
    KeyUri = new Uri("https://contosovault.vault.azure.net/keys/azureKey"),
    MobileNetworkId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
};
ArmOperation<MobileNetworkSimGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, simGroupName, data);
MobileNetworkSimGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileNetworkSimGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
