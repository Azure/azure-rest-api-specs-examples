using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectedItems_SwitchProvider.json
// this example is just showing the usage of "ReplicationProtectedItems_SwitchProvider" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
SwitchProviderContent content = new SwitchProviderContent
{
    Properties = new SwitchProviderProperties
    {
        TargetInstanceType = "InMageRcm",
        ProviderSpecificDetails = new InMageAzureV2SwitchProviderContent(new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault2"), new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud2"), "5efaa202-e958-435e-8171-706bf735fcc4"),
    },
};
ArmOperation<ReplicationProtectedItemResource> lro = await replicationProtectedItem.SwitchProviderAsync(WaitUntil.Completed, content);
ReplicationProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReplicationProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
