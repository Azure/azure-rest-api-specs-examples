using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/DataVersionBase/registryList.json
// this example is just showing the usage of "RegistryDataVersions_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
string orderBy = "string";
int? top = 1;
string tags = "string";
MachineLearningListViewType? listViewType = MachineLearningListViewType.ArchivedOnly;
await foreach (MachineLearningRegistryDataVersionResource item in collection.GetAllAsync(orderBy: orderBy, top: top, tags: tags, listViewType: listViewType))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningDataVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
