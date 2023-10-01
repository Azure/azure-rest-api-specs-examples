using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Registries/update.json
// this example is just showing the usage of "Registries_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryResource created on azure
// for more information of creating MachineLearningRegistryResource, please refer to the document of MachineLearningRegistryResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "string";
ResourceIdentifier machineLearningRegistryResourceId = MachineLearningRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
MachineLearningRegistryResource machineLearningRegistry = client.GetMachineLearningRegistryResource(machineLearningRegistryResourceId);

// invoke the operation
MachineLearningRegistryPatch patch = new MachineLearningRegistryPatch()
{
    Identity = new ManagedServiceIdentity("SystemAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("string")] = new UserAssignedIdentity(),
        },
    },
    Sku = new MachineLearningSkuPatch()
    {
        Capacity = 1,
        Family = "string",
        Name = "string",
        Size = "string",
        Tier = MachineLearningSkuTier.Basic,
    },
    Tags =
    {
    },
};
MachineLearningRegistryResource result = await machineLearningRegistry.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningRegistryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
