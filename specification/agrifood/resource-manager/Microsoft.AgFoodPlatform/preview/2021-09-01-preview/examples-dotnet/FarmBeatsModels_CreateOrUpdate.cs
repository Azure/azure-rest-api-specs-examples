using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AgFoodPlatform;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_CreateOrUpdate.json
// this example is just showing the usage of "FarmBeatsModels_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "examples-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this FarmBeatResource
FarmBeatCollection collection = resourceGroupResource.GetFarmBeats();

// invoke the operation
string farmBeatsResourceName = "examples-farmbeatsResourceName";
FarmBeatData data = new FarmBeatData(new AzureLocation("eastus2"))
{
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2",
    },
};
ArmOperation<FarmBeatResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, farmBeatsResourceName, data);
FarmBeatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FarmBeatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
