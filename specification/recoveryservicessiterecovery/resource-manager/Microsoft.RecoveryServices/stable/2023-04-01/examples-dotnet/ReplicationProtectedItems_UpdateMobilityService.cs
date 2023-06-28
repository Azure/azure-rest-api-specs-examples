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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ReplicationProtectedItems_UpdateMobilityService.json
// this example is just showing the usage of "ReplicationProtectedItems_UpdateMobilityService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReplicationProtectedItemResource created on azure
// for more information of creating ReplicationProtectedItemResource, please refer to the document of ReplicationProtectedItemResource
string subscriptionId = "b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c";
string resourceGroupName = "wcusValidations";
string resourceName = "WCUSVault";
string fabricName = "WIN-JKKJ31QI8U2";
string protectionContainerName = "cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0";
string replicatedProtectedItemName = "79dd20ab-2b40-11e7-9791-0050568f387e";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// invoke the operation
UpdateMobilityServiceContent content = new UpdateMobilityServiceContent()
{
    UpdateMobilityServiceRequestRunAsAccountId = "2",
};
ArmOperation<ReplicationProtectedItemResource> lro = await replicationProtectedItem.UpdateMobilityServiceAsync(WaitUntil.Completed, content);
ReplicationProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReplicationProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
