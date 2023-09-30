using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/FeaturestoreEntityVersion/list.json
// this example is just showing the usage of "FeaturestoreEntityVersions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningFeatureStoreEntityContainerResource created on azure
// for more information of creating MachineLearningFeatureStoreEntityContainerResource, please refer to the document of MachineLearningFeatureStoreEntityContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningFeatureStoreEntityContainerResourceId = MachineLearningFeatureStoreEntityContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningFeatureStoreEntityContainerResource machineLearningFeatureStoreEntityContainer = client.GetMachineLearningFeatureStoreEntityContainerResource(machineLearningFeatureStoreEntityContainerResourceId);

// get the collection of this MachineLearningFeaturestoreEntityVersionResource
MachineLearningFeaturestoreEntityVersionCollection collection = machineLearningFeatureStoreEntityContainer.GetMachineLearningFeaturestoreEntityVersions();

// invoke the operation and iterate over the result
MachineLearningFeaturestoreEntityVersionCollectionGetAllOptions options = new MachineLearningFeaturestoreEntityVersionCollectionGetAllOptions() { Tags = "string", ListViewType = MachineLearningListViewType.ActiveOnly };
await foreach (MachineLearningFeaturestoreEntityVersionResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningFeaturestoreEntityVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
