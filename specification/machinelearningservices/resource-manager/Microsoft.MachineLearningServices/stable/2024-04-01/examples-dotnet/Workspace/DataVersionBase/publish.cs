using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/DataVersionBase/publish.json
// this example is just showing the usage of "DataVersions_Publish" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningDataVersionResource created on azure
// for more information of creating MachineLearningDataVersionResource, please refer to the document of MachineLearningDataVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningDataVersionResourceId = MachineLearningDataVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name, version);
MachineLearningDataVersionResource machineLearningDataVersion = client.GetMachineLearningDataVersionResource(machineLearningDataVersionResourceId);

// invoke the operation
DestinationAssetContent content = new DestinationAssetContent
{
    RegistryName = "string",
    DestinationName = "string",
    DestinationVersion = "string",
};
await machineLearningDataVersion.PublishAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
