using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Registry/ModelVersion/get.json
// this example is just showing the usage of "RegistryModelVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryModelContainerResource created on azure
// for more information of creating MachineLearningRegistryModelContainerResource, please refer to the document of MachineLearningRegistryModelContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "my-aml-registry";
string modelName = "string";
ResourceIdentifier machineLearningRegistryModelContainerResourceId = MachineLearningRegistryModelContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, modelName);
MachineLearningRegistryModelContainerResource machineLearningRegistryModelContainer = client.GetMachineLearningRegistryModelContainerResource(machineLearningRegistryModelContainerResourceId);

// get the collection of this MachineLearningRegistryModelVersionResource
MachineLearningRegistryModelVersionCollection collection = machineLearningRegistryModelContainer.GetMachineLearningRegistryModelVersions();

// invoke the operation
string version = "string";
NullableResponse<MachineLearningRegistryModelVersionResource> response = await collection.GetIfExistsAsync(version);
MachineLearningRegistryModelVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningModelVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
