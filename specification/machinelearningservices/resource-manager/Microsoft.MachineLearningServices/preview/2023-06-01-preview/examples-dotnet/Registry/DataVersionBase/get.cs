using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Registry/DataVersionBase/get.json
// this example is just showing the usage of "RegistryDataVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryDataContainerResource created on azure
// for more information of creating MachineLearningRegistryDataContainerResource, please refer to the document of MachineLearningRegistryDataContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "registryName";
string name = "string";
ResourceIdentifier machineLearningRegistryDataContainerResourceId = MachineLearningRegistryDataContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, name);
MachineLearningRegistryDataContainerResource machineLearningRegistryDataContainer = client.GetMachineLearningRegistryDataContainerResource(machineLearningRegistryDataContainerResourceId);

// get the collection of this MachineLearningRegistryDataVersionResource
MachineLearningRegistryDataVersionCollection collection = machineLearningRegistryDataContainer.GetMachineLearningRegistryDataVersions();

// invoke the operation
string version = "string";
NullableResponse<MachineLearningRegistryDataVersionResource> response = await collection.GetIfExistsAsync(version);
MachineLearningRegistryDataVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningDataVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
