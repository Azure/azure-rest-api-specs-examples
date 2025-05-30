using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ClusterRecoveryPoint_Get.json
// this example is just showing the usage of "ClusterRecoveryPoint_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryReplicationProtectionClusterResource created on azure
// for more information of creating SiteRecoveryReplicationProtectionClusterResource, please refer to the document of SiteRecoveryReplicationProtectionClusterResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "fabric-pri-eastus";
string protectionContainerName = "pri-cloud-eastus";
string replicationProtectionClusterName = "testcluster";
ResourceIdentifier siteRecoveryReplicationProtectionClusterResourceId = SiteRecoveryReplicationProtectionClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicationProtectionClusterName);
SiteRecoveryReplicationProtectionClusterResource siteRecoveryReplicationProtectionClusterResource = client.GetSiteRecoveryReplicationProtectionClusterResource(siteRecoveryReplicationProtectionClusterResourceId);

// get the collection of this SiteRecoveryClusterRecoveryPointResource
SiteRecoveryClusterRecoveryPointCollection collection = siteRecoveryReplicationProtectionClusterResource.GetSiteRecoveryClusterRecoveryPoints();

// invoke the operation
string recoveryPointName = "06b9ae7f-f21d-4a76-9897-5cf5d6004d80";
NullableResponse<SiteRecoveryClusterRecoveryPointResource> response = await collection.GetIfExistsAsync(recoveryPointName);
SiteRecoveryClusterRecoveryPointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SiteRecoveryClusterRecoveryPointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
