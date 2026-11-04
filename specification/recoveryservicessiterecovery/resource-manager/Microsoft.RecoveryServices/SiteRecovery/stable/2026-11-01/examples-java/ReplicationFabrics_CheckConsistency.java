
/**
 * Samples for ReplicationFabrics CheckConsistency.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationFabrics_CheckConsistency.json
     */
    /**
     * Sample code: Checks the consistency of the ASR fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void checksTheConsistencyOfTheASRFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().checkConsistency("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
