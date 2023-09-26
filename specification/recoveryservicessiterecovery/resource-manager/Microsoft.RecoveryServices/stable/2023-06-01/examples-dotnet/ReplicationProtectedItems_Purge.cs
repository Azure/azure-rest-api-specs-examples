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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectedItems_Purge.json
// this example is just showing the usage of "ReplicationProtectedItems_Purge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReplicationProtectedItemResource created on azure
// for more information of creating ReplicationProtectedItemResource, please refer to the document of ReplicationProtectedItemResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "cloud1";
string protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
string replicatedProtectedItemName = "c0c14913-3d7a-48ea-9531-cc99e0e686e6";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// invoke the operation
await replicationProtectedItem.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
