using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/DataVersionBase/get.json
// this example is just showing the usage of "DataVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningDataContainerResource created on azure
// for more information of creating MachineLearningDataContainerResource, please refer to the document of MachineLearningDataContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningDataContainerResourceId = MachineLearningDataContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningDataContainerResource machineLearningDataContainer = client.GetMachineLearningDataContainerResource(machineLearningDataContainerResourceId);

// get the collection of this MachineLearningDataVersionResource
MachineLearningDataVersionCollection collection = machineLearningDataContainer.GetMachineLearningDataVersions();

// invoke the operation
string version = "string";
bool result = await collection.ExistsAsync(version);

Console.WriteLine($"Succeeded: {result}");
