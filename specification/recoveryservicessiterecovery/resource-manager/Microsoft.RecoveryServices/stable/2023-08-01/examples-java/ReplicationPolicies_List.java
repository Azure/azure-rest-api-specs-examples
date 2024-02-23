
/**
 * Samples for ReplicationPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationPolicies_List.json
     */
    /**
     * Sample code: Gets the list of replication policies.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationPolicies(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationPolicies().list("vault1", "resourceGroupPS1", com.azure.core.util.Context.NONE);
    }
}
