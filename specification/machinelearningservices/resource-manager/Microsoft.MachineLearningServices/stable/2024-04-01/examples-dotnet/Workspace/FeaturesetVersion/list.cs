using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/FeaturesetVersion/list.json
// this example is just showing the usage of "FeaturesetVersions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningFeatureSetContainerResource created on azure
// for more information of creating MachineLearningFeatureSetContainerResource, please refer to the document of MachineLearningFeatureSetContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningFeatureSetContainerResourceId = MachineLearningFeatureSetContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningFeatureSetContainerResource machineLearningFeatureSetContainer = client.GetMachineLearningFeatureSetContainerResource(machineLearningFeatureSetContainerResourceId);

// get the collection of this MachineLearningFeatureSetVersionResource
MachineLearningFeatureSetVersionCollection collection = machineLearningFeatureSetContainer.GetMachineLearningFeatureSetVersions();

// invoke the operation and iterate over the result
MachineLearningFeatureSetVersionCollectionGetAllOptions options = new MachineLearningFeatureSetVersionCollectionGetAllOptions() { Tags = "string", ListViewType = MachineLearningListViewType.All };
await foreach (MachineLearningFeatureSetVersionResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningFeatureSetVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
