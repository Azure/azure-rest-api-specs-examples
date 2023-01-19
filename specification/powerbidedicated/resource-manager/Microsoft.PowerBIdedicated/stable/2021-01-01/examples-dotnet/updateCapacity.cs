using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PowerBIDedicated;
using Azure.ResourceManager.PowerBIDedicated.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateCapacity.json
// this example is just showing the usage of "Capacities_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedCapacityResource created on azure
// for more information of creating DedicatedCapacityResource, please refer to the document of DedicatedCapacityResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
string dedicatedCapacityName = "azsdktest";
ResourceIdentifier dedicatedCapacityResourceId = DedicatedCapacityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dedicatedCapacityName);
DedicatedCapacityResource dedicatedCapacity = client.GetDedicatedCapacityResource(dedicatedCapacityResourceId);

// invoke the operation
DedicatedCapacityPatch patch = new DedicatedCapacityPatch()
{
    Sku = new CapacitySku("A1")
    {
        Tier = CapacitySkuTier.PbieAzure,
    },
    Tags =
    {
    ["testKey"] = "testValue",
    },
    AdministrationMembers =
    {
    "azsdktest@microsoft.com","azsdktest2@microsoft.com"
    },
};
ArmOperation<DedicatedCapacityResource> lro = await dedicatedCapacity.UpdateAsync(WaitUntil.Completed, patch);
DedicatedCapacityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedCapacityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
