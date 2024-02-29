using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/Pools_Get.json
// this example is just showing the usage of "Pools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppAccountResource created on azure
// for more information of creating NetAppAccountResource, please refer to the document of NetAppAccountResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
ResourceIdentifier netAppAccountResourceId = NetAppAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
NetAppAccountResource netAppAccount = client.GetNetAppAccountResource(netAppAccountResourceId);

// get the collection of this CapacityPoolResource
CapacityPoolCollection collection = netAppAccount.GetCapacityPools();

// invoke the operation
string poolName = "pool1";
NullableResponse<CapacityPoolResource> response = await collection.GetIfExistsAsync(poolName);
CapacityPoolResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CapacityPoolData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
