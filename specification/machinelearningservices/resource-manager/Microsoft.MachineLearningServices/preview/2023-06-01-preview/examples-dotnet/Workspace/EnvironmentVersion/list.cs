using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/EnvironmentVersion/list.json
// this example is just showing the usage of "EnvironmentVersions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningEnvironmentContainerResource created on azure
// for more information of creating MachineLearningEnvironmentContainerResource, please refer to the document of MachineLearningEnvironmentContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningEnvironmentContainerResourceId = MachineLearningEnvironmentContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningEnvironmentContainerResource machineLearningEnvironmentContainer = client.GetMachineLearningEnvironmentContainerResource(machineLearningEnvironmentContainerResourceId);

// get the collection of this MachineLearningEnvironmentVersionResource
MachineLearningEnvironmentVersionCollection collection = machineLearningEnvironmentContainer.GetMachineLearningEnvironmentVersions();

// invoke the operation and iterate over the result
string orderBy = "string";
int? top = 1;
await foreach (MachineLearningEnvironmentVersionResource item in collection.GetAllAsync(orderBy: orderBy, top: top))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningEnvironmentVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
