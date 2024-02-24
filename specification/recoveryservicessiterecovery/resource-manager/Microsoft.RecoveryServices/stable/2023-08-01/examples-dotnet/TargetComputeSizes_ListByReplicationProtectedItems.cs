using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/TargetComputeSizes_ListByReplicationProtectedItems.json
// this example is just showing the usage of "TargetComputeSizes_ListByReplicationProtectedItems" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReplicationProtectedItemResource created on azure
// for more information of creating ReplicationProtectedItemResource, please refer to the document of ReplicationProtectedItemResource
string subscriptionId = "6808dbbc-98c7-431f-a1b1-9580902423b7";
string resourceGroupName = "avraiMgDiskVaultRG";
string resourceName = "avraiMgDiskVault";
string fabricName = "asr-a2a-default-centraluseuap";
string protectionContainerName = "asr-a2a-default-centraluseuap-container";
string replicatedProtectedItemName = "468c912d-b1ab-4ea2-97eb-4b5095155db2";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// invoke the operation and iterate over the result
await foreach (TargetComputeSize item in replicationProtectedItem.GetTargetComputeSizesByReplicationProtectedItemsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
