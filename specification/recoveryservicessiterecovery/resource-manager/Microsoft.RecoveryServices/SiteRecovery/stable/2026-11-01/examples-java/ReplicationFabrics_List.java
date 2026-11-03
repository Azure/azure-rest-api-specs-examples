
/**
 * Samples for ReplicationFabrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationFabrics_List.json
     */
    /**
     * Sample code: Gets the list of ASR fabrics.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheListOfASRFabrics(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
