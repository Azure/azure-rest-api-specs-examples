using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/Workflows_Abort.json
// this example is just showing the usage of "Workflows_Abort" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncWorkflowResource created on azure
// for more information of creating StorageSyncWorkflowResource, please refer to the document of StorageSyncWorkflowResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
string workflowId = "7ffd50b3-5574-478d-9ff2-9371bc42ce68";
ResourceIdentifier storageSyncWorkflowResourceId = StorageSyncWorkflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, workflowId);
StorageSyncWorkflowResource storageSyncWorkflow = client.GetStorageSyncWorkflowResource(storageSyncWorkflowResourceId);

// invoke the operation
await storageSyncWorkflow.AbortAsync();

Console.WriteLine($"Succeeded");
