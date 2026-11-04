
/**
 * Samples for ReplicationProtectionClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_List.json
     */
    /**
     * Sample code: Gets the list of Replication protection clusters in vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectionClustersInVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().list("resourceGroupPS1", "vault1", null,
            "SourceFabricName eq 'asr-a2a-default-eastus' and SourceFabricLocation eq 'East US' and InstanceType eq 'A2A'",
            com.azure.core.util.Context.NONE);
    }
}
