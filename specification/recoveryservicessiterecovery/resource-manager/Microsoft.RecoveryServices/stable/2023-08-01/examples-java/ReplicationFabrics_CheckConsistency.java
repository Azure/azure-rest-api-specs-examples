
/**
 * Samples for ReplicationFabrics CheckConsistency.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationFabrics_CheckConsistency.json
     */
    /**
     * Sample code: Checks the consistency of the ASR fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void checksTheConsistencyOfTheASRFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().checkConsistency("vault1", "resourceGroupPS1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
