using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/CodeVersion/createOrGetStartPendingUpload.json
// this example is just showing the usage of "RegistryCodeVersions_CreateOrGetStartPendingUpload" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryCodeVersionResource created on azure
// for more information of creating MachineLearningRegistryCodeVersionResource, please refer to the document of MachineLearningRegistryCodeVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "registryName";
string codeName = "string";
string version = "string";
ResourceIdentifier machineLearningRegistryCodeVersionResourceId = MachineLearningRegistryCodeVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, codeName, version);
MachineLearningRegistryCodeVersionResource machineLearningRegistryCodeVersion = client.GetMachineLearningRegistryCodeVersionResource(machineLearningRegistryCodeVersionResourceId);

// invoke the operation
PendingUploadRequestDto body = new PendingUploadRequestDto()
{
    PendingUploadId = "string",
    PendingUploadType = PendingUploadType.TemporaryBlobReference,
};
PendingUploadResponseDto result = await machineLearningRegistryCodeVersion.CreateOrGetStartPendingUploadAsync(body);

Console.WriteLine($"Succeeded: {result}");
