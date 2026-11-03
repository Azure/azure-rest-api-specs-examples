
/**
 * Samples for ReplicationPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationPolicies_List.json
     */
    /**
     * Sample code: Gets the list of replication policies.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationPolicies(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationPolicies().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
