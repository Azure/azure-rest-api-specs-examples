using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationStorageClassificationMappings_Delete.json
// this example is just showing the usage of "ReplicationStorageClassificationMappings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageClassificationMappingResource created on azure
// for more information of creating StorageClassificationMappingResource, please refer to the document of StorageClassificationMappingResource
string subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0";
string storageClassificationName = "8891569e-aaef-4a46-a4a0-78c14f2d7b09";
string storageClassificationMappingName = "testStorageMapping";
ResourceIdentifier storageClassificationMappingResourceId = StorageClassificationMappingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, storageClassificationName, storageClassificationMappingName);
StorageClassificationMappingResource storageClassificationMapping = client.GetStorageClassificationMappingResource(storageClassificationMappingResourceId);

// invoke the operation
await storageClassificationMapping.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
