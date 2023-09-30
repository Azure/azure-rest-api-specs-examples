using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/EnvironmentVersion/get.json
// this example is just showing the usage of "EnvironmentVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningEnvironmentVersionResource created on azure
// for more information of creating MachineLearningEnvironmentVersionResource, please refer to the document of MachineLearningEnvironmentVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningEnvironmentVersionResourceId = MachineLearningEnvironmentVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name, version);
MachineLearningEnvironmentVersionResource machineLearningEnvironmentVersion = client.GetMachineLearningEnvironmentVersionResource(machineLearningEnvironmentVersionResourceId);

// invoke the operation
MachineLearningEnvironmentVersionResource result = await machineLearningEnvironmentVersion.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningEnvironmentVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
