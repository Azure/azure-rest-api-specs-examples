using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/RecoveryPoints_Get.json
// this example is just showing the usage of "RecoveryPoints_Get" operation, for the dependent resources, they will have to be created separately.

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
string replicatedProtectedItemName = "f8491e4f-817a-40dd-a90c-af773978c75b";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// get the collection of this SiteRecoveryPointResource
SiteRecoveryPointCollection collection = replicationProtectedItem.GetSiteRecoveryPoints();

// invoke the operation
string recoveryPointName = "b22134ea-620c-474b-9fa5-3c1cb47708e3";
bool result = await collection.ExistsAsync(recoveryPointName);

Console.WriteLine($"Succeeded: {result}");
