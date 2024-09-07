using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/CodeVersion/list.json
// this example is just showing the usage of "RegistryCodeVersions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryCodeContainerResource created on azure
// for more information of creating MachineLearningRegistryCodeContainerResource, please refer to the document of MachineLearningRegistryCodeContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "my-aml-registry";
string codeName = "string";
ResourceIdentifier machineLearningRegistryCodeContainerResourceId = MachineLearningRegistryCodeContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, codeName);
MachineLearningRegistryCodeContainerResource machineLearningRegistryCodeContainer = client.GetMachineLearningRegistryCodeContainerResource(machineLearningRegistryCodeContainerResourceId);

// get the collection of this MachineLearningRegistryCodeVersionResource
MachineLearningRegistryCodeVersionCollection collection = machineLearningRegistryCodeContainer.GetMachineLearningRegistryCodeVersions();

// invoke the operation and iterate over the result
string orderBy = "string";
int? top = 1;
await foreach (MachineLearningRegistryCodeVersionResource item in collection.GetAllAsync(orderBy: orderBy, top: top))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningCodeVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
