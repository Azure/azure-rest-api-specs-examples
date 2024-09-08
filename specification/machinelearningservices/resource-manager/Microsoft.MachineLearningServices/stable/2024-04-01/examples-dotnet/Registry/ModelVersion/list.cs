using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/ModelVersion/list.json
// this example is just showing the usage of "RegistryModelVersions_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
MachineLearningRegistryModelVersionCollectionGetAllOptions options = new MachineLearningRegistryModelVersionCollectionGetAllOptions() { OrderBy = "string", Top = 1, Version = "string", Description = "string", Tags = "string", Properties = "string" };
await foreach (MachineLearningRegistryModelVersionResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningModelVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
