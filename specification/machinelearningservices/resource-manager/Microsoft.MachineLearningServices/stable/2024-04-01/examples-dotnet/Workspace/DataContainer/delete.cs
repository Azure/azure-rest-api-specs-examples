using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/DataContainer/delete.json
// this example is just showing the usage of "DataContainers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningDataContainerResource created on azure
// for more information of creating MachineLearningDataContainerResource, please refer to the document of MachineLearningDataContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string workspaceName = "workspace123";
string name = "datacontainer123";
ResourceIdentifier machineLearningDataContainerResourceId = MachineLearningDataContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningDataContainerResource machineLearningDataContainer = client.GetMachineLearningDataContainerResource(machineLearningDataContainerResourceId);

// invoke the operation
await machineLearningDataContainer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
