using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PowerBIDedicated;
using Azure.ResourceManager.PowerBIDedicated.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createAutoScaleVCore.json
// this example is just showing the usage of "AutoScaleVCores_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AutoScaleVCoreResource
AutoScaleVCoreCollection collection = resourceGroupResource.GetAutoScaleVCores();

// invoke the operation
string vcoreName = "testvcore";
AutoScaleVCoreData data = new AutoScaleVCoreData(new AzureLocation("West US"), new AutoScaleVCoreSku("AutoScale")
{
    Tier = VCoreSkuTier.AutoScale,
    Capacity = 0,
})
{
    CapacityLimit = 10,
    CapacityObjectId = "a28f00bd-5330-4572-88f1-fa883e074785",
    Tags =
    {
    ["testKey"] = "testValue",
    },
};
ArmOperation<AutoScaleVCoreResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vcoreName, data);
AutoScaleVCoreResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutoScaleVCoreData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
