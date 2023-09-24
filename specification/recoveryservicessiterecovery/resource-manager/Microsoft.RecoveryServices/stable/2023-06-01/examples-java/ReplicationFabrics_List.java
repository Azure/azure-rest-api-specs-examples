/** Samples for ReplicationFabrics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationFabrics_List.json
     */
    /**
     * Sample code: Gets the list of ASR fabrics.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfASRFabrics(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().list("vault1", "resourceGroupPS1", com.azure.core.util.Context.NONE);
    }
}
