using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutLogicalNetwork.json
// this example is just showing the usage of "LogicalNetworks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this LogicalNetworkResource
LogicalNetworkCollection collection = resourceGroupResource.GetLogicalNetworks();

// invoke the operation
string logicalNetworkName = "test-lnet";
LogicalNetworkData data = new LogicalNetworkData(new AzureLocation("West US2"))
{
    ExtendedLocation = new ArcVmExtendedLocation()
    {
        Name = "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
        ExtendedLocationType = ArcVmExtendedLocationType.CustomLocation,
    },
};
ArmOperation<LogicalNetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, logicalNetworkName, data);
LogicalNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogicalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
