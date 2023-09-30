using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Registry/DataVersionBase/createOrGetStartPendingUpload.json
// this example is just showing the usage of "RegistryDataVersions_CreateOrGetStartPendingUpload" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryDataVersionResource created on azure
// for more information of creating MachineLearningRegistryDataVersionResource, please refer to the document of MachineLearningRegistryDataVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "registryName";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningRegistryDataVersionResourceId = MachineLearningRegistryDataVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, name, version);
MachineLearningRegistryDataVersionResource machineLearningRegistryDataVersion = client.GetMachineLearningRegistryDataVersionResource(machineLearningRegistryDataVersionResourceId);

// invoke the operation
PendingUploadRequestDto body = new PendingUploadRequestDto()
{
    PendingUploadId = "string",
    PendingUploadType = PendingUploadType.None,
};
PendingUploadResponseDto result = await machineLearningRegistryDataVersion.CreateOrGetStartPendingUploadAsync(body);

Console.WriteLine($"Succeeded: {result}");
