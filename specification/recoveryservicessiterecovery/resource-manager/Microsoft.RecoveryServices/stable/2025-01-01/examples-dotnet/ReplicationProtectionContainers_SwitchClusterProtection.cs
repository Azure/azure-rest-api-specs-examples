using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionContainers_SwitchClusterProtection.json
// this example is just showing the usage of "ReplicationProtectionContainers_SwitchClusterProtection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryProtectionContainerResource created on azure
// for more information of creating SiteRecoveryProtectionContainerResource, please refer to the document of SiteRecoveryProtectionContainerResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "fabric-pri-eastus";
string protectionContainerName = "pri-cloud-eastus";
ResourceIdentifier siteRecoveryProtectionContainerResourceId = SiteRecoveryProtectionContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName);
SiteRecoveryProtectionContainerResource siteRecoveryProtectionContainer = client.GetSiteRecoveryProtectionContainerResource(siteRecoveryProtectionContainerResourceId);

// invoke the operation
SwitchClusterProtectionContent content = new SwitchClusterProtectionContent
{
    Properties = new SwitchClusterProtectionContentProperties
    {
        ReplicationProtectionClusterName = "testcluster",
        ProviderSpecificDetails = new A2ASwitchClusterProtectionContent
        {
            RecoveryContainerId = new ResourceIdentifier("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus/replicationProtectionContainers/rec-cloud-westus"),
            PolicyId = new ResourceIdentifier("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/klncksan"),
            ProtectedItemsDetail = {new A2AProtectedItemDetail
            {
            VmManagedDisks = {new A2AVmManagedDiskDetails("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql0-osdisk", new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache"), new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"))},
            RecoveryResourceGroupId = new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
            ReplicationProtectedItemName = "yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio",
            }, new A2AProtectedItemDetail
            {
            VmManagedDisks = {new A2AVmManagedDiskDetails("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql1-osdisk", new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache"), new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"))},
            RecoveryResourceGroupId = new ResourceIdentifier("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
            ReplicationProtectedItemName = "kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4",
            }},
        },
    },
};
ArmOperation<SiteRecoveryProtectionContainerResource> lro = await siteRecoveryProtectionContainer.SwitchClusterProtectionAsync(WaitUntil.Completed, content);
SiteRecoveryProtectionContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryProtectionContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
