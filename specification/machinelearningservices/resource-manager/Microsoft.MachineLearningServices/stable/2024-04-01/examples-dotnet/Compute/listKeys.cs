using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/listKeys.json
// this example is just showing the usage of "Compute_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningComputeResource created on azure
// for more information of creating MachineLearningComputeResource, please refer to the document of MachineLearningComputeResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string workspaceName = "workspaces123";
string computeName = "compute123";
ResourceIdentifier machineLearningComputeResourceId = MachineLearningComputeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, computeName);
MachineLearningComputeResource machineLearningCompute = client.GetMachineLearningComputeResource(machineLearningComputeResourceId);

// invoke the operation
MachineLearningComputeSecrets result = await machineLearningCompute.GetKeysAsync();

Console.WriteLine($"Succeeded: {result}");
