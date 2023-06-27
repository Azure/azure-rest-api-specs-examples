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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ReplicationProtectedItems_UpdateAppliance.json
// this example is just showing the usage of "ReplicationProtectedItems_UpdateAppliance" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReplicationProtectedItemResource created on azure
// for more information of creating ReplicationProtectedItemResource, please refer to the document of ReplicationProtectedItemResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "Ayan-0106-SA-RG";
string resourceName = "Ayan-0106-SA-Vault";
string fabricName = "Ayan-0106-SA-Vaultreplicationfabric";
string protectionContainerName = "Ayan-0106-SA-Vaultreplicationcontainer";
string replicatedProtectedItemName = "idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// invoke the operation
UpdateApplianceForReplicationProtectedItemContent content = new UpdateApplianceForReplicationProtectedItemContent(new UpdateApplianceForReplicationProtectedItemProperties("", new InMageRcmUpdateApplianceForReplicationProtectedItemContent()
{
    RunAsAccountId = "",
}));
ArmOperation<ReplicationProtectedItemResource> lro = await replicationProtectedItem.UpdateApplianceAsync(WaitUntil.Completed, content);
ReplicationProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReplicationProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
