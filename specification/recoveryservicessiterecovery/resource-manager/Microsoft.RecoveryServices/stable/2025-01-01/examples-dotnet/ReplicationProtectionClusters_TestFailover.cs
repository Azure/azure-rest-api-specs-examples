using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionClusters_TestFailover.json
// this example is just showing the usage of "ReplicationProtectionClusters_TestFailover" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ClusterTestFailoverContent content = new ClusterTestFailoverContent(new ClusterTestFailoverContentProperties
{
    FailoverDirection = RecoveryServicesSiteRecoveryFailoverDirection.PrimaryToRecovery,
    NetworkType = "VmNetworkAsInput",
    NetworkId = new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr/providers/Microsoft.Network/virtualNetworks/adVNET-asr"),
    ProviderSpecificDetails = new A2AClusterTestFailoverContent
    {
        ClusterRecoveryPointId = new ResourceIdentifier("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/cc48b7f3-b267-432b-ad76-45528974dc62"),
        IndividualNodeRecoveryPoints = { "/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/testVM/recoveryPoints/b5c2051b-79e3-41ad-9d25-244f6ef8ce7d" },
    },
});
ArmOperation<SiteRecoveryReplicationProtectionClusterResource> lro = await siteRecoveryReplicationProtectionClusterResource.TestFailoverAsync(WaitUntil.Completed, content);
SiteRecoveryReplicationProtectionClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryReplicationProtectionClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
