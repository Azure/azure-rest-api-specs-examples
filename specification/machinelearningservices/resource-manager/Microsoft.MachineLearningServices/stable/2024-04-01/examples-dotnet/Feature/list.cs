using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Feature/list.json
// this example is just showing the usage of "Features_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningFeatureSetVersionResource created on azure
// for more information of creating MachineLearningFeatureSetVersionResource, please refer to the document of MachineLearningFeatureSetVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string featuresetName = "string";
string featuresetVersion = "string";
ResourceIdentifier machineLearningFeatureSetVersionResourceId = MachineLearningFeatureSetVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, featuresetName, featuresetVersion);
MachineLearningFeatureSetVersionResource machineLearningFeatureSetVersion = client.GetMachineLearningFeatureSetVersionResource(machineLearningFeatureSetVersionResourceId);

// get the collection of this MachineLearningFeatureResource
MachineLearningFeatureCollection collection = machineLearningFeatureSetVersion.GetMachineLearningFeatures();

// invoke the operation and iterate over the result
MachineLearningFeatureCollectionGetAllOptions options = new MachineLearningFeatureCollectionGetAllOptions { Tags = "string", FeatureName = "string", Description = "string" };
await foreach (MachineLearningFeatureResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningFeatureData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
