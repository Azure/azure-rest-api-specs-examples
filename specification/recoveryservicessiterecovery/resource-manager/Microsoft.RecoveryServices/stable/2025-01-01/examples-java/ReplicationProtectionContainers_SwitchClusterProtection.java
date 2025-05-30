
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AProtectedItemDetail;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2ASwitchClusterProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AVmManagedDiskInputDetails;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchClusterProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchClusterProtectionInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectionContainers SwitchClusterProtection.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainers_SwitchClusterProtection.json
     */
    /**
     * Sample code: Switches protection from one container to another or one replication provider to another.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void switchesProtectionFromOneContainerToAnotherOrOneReplicationProviderToAnother(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().switchClusterProtection("resourceGroupPS1", "vault1",
            "fabric-pri-eastus", "pri-cloud-eastus",
            new SwitchClusterProtectionInput().withProperties(new SwitchClusterProtectionInputProperties()
                .withReplicationProtectionClusterName("testcluster")
                .withProviderSpecificDetails(new A2ASwitchClusterProtectionInput().withRecoveryContainerId(
                    "/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus/replicationProtectionContainers/rec-cloud-westus")
                    .withPolicyId(
                        "/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/klncksan")
                    .withProtectedItemsDetail(Arrays.asList(new A2AProtectedItemDetail()
                        .withVmManagedDisks(Arrays.asList(new A2AVmManagedDiskInputDetails().withDiskId(
                            "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql0-osdisk")
                            .withPrimaryStagingAzureStorageAccountId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache")
                            .withRecoveryResourceGroupId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr")))
                        .withRecoveryResourceGroupId(
                            "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr")
                        .withReplicationProtectedItemName("yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio"),
                        new A2AProtectedItemDetail().withVmManagedDisks(Arrays.asList(new A2AVmManagedDiskInputDetails()
                            .withDiskId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql1-osdisk")
                            .withPrimaryStagingAzureStorageAccountId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache")
                            .withRecoveryResourceGroupId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr")))
                            .withRecoveryResourceGroupId(
                                "/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr")
                            .withReplicationProtectedItemName("kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4"))))),
            com.azure.core.util.Context.NONE);
    }
}
