using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/EnvironmentContainer/delete.json
// this example is just showing the usage of "EnvironmentContainers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningEnvironmentContainerResource created on azure
// for more information of creating MachineLearningEnvironmentContainerResource, please refer to the document of MachineLearningEnvironmentContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string workspaceName = "testworkspace";
string name = "testContainer";
ResourceIdentifier machineLearningEnvironmentContainerResourceId = MachineLearningEnvironmentContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningEnvironmentContainerResource machineLearningEnvironmentContainer = client.GetMachineLearningEnvironmentContainerResource(machineLearningEnvironmentContainerResourceId);

// invoke the operation
await machineLearningEnvironmentContainer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
