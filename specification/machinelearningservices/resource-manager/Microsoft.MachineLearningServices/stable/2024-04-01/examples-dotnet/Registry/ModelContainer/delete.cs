using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/ModelContainer/delete.json
// this example is just showing the usage of "RegistryModelContainers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryModelContainerResource created on azure
// for more information of creating MachineLearningRegistryModelContainerResource, please refer to the document of MachineLearningRegistryModelContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string registryName = "registry123";
string modelName = "testContainer";
ResourceIdentifier machineLearningRegistryModelContainerResourceId = MachineLearningRegistryModelContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, modelName);
MachineLearningRegistryModelContainerResource machineLearningRegistryModelContainer = client.GetMachineLearningRegistryModelContainerResource(machineLearningRegistryModelContainerResourceId);

// invoke the operation
await machineLearningRegistryModelContainer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
