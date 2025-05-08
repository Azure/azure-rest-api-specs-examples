using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-01-01/examples/Pools_CreateOrUpdate.json
// this example is just showing the usage of "Pools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppAccountResource created on azure
// for more information of creating NetAppAccountResource, please refer to the document of NetAppAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
ResourceIdentifier netAppAccountResourceId = NetAppAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
NetAppAccountResource netAppAccount = client.GetNetAppAccountResource(netAppAccountResourceId);

// get the collection of this CapacityPoolResource
CapacityPoolCollection collection = netAppAccount.GetCapacityPools();

// invoke the operation
string poolName = "pool1";
CapacityPoolData data = new CapacityPoolData(new AzureLocation("eastus"), 4398046511104L, NetAppFileServiceLevel.Premium)
{
    QosType = CapacityPoolQosType.Auto,
};
ArmOperation<CapacityPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
CapacityPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CapacityPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
