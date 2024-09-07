using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/ModelVersion/list.json
// this example is just showing the usage of "ModelVersions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningModelContainerResource created on azure
// for more information of creating MachineLearningModelContainerResource, please refer to the document of MachineLearningModelContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningModelContainerResourceId = MachineLearningModelContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningModelContainerResource machineLearningModelContainer = client.GetMachineLearningModelContainerResource(machineLearningModelContainerResourceId);

// get the collection of this MachineLearningModelVersionResource
MachineLearningModelVersionCollection collection = machineLearningModelContainer.GetMachineLearningModelVersions();

// invoke the operation and iterate over the result
MachineLearningModelVersionCollectionGetAllOptions options = new MachineLearningModelVersionCollectionGetAllOptions() { OrderBy = "string", Top = 1, Version = "string", Description = "string", Offset = 1, Tags = "string", Properties = "string" };
await foreach (MachineLearningModelVersionResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningModelVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
