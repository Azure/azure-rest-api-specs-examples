using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Fabric.Models;
using Azure.ResourceManager.Fabric;

// Generated from example definition: 2023-11-01/FabricCapacities_Update.json
// this example is just showing the usage of "FabricCapacity_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FabricCapacityResource created on azure
// for more information of creating FabricCapacityResource, please refer to the document of FabricCapacityResource
string subscriptionId = "548B7FB7-3B2A-4F46-BB02-66473F1FC22C";
string resourceGroupName = "TestRG";
string capacityName = "azsdktest";
ResourceIdentifier fabricCapacityResourceId = FabricCapacityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityName);
FabricCapacityResource fabricCapacity = client.GetFabricCapacityResource(fabricCapacityResourceId);

// invoke the operation
FabricCapacityPatch patch = new FabricCapacityPatch
{
    Sku = new FabricSku("F8", FabricSkuTier.Fabric),
    Tags =
    {
    ["testKey"] = "testValue"
    },
    AdministrationMembers = { "azsdktest2@microsoft.com" },
};
ArmOperation<FabricCapacityResource> lro = await fabricCapacity.UpdateAsync(WaitUntil.Completed, patch);
FabricCapacityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FabricCapacityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
