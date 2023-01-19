using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AgFoodPlatform;
using Azure.ResourceManager.AgFoodPlatform.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_GetOperationResult.json
// this example is just showing the usage of "FarmBeatsModels_GetOperationResult" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FarmBeatResource created on azure
// for more information of creating FarmBeatResource, please refer to the document of FarmBeatResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "examples-rg";
string farmBeatsResourceName = "examples-farmBeatsResourceName";
ResourceIdentifier farmBeatResourceId = FarmBeatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, farmBeatsResourceName);
FarmBeatResource farmBeat = client.GetFarmBeatResource(farmBeatResourceId);

// invoke the operation
string operationResultsId = "resource-provisioning-id-farmBeatsResourceName";
ArmAsyncOperation result = await farmBeat.GetOperationResultAsync(operationResultsId);

Console.WriteLine($"Succeeded: {result}");
