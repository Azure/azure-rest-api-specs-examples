
/**
 * Samples for ReplicationEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationEvents_List.json
     */
    /**
     * Sample code: Gets the list of Azure Site Recovery events.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfAzureSiteRecoveryEvents(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationEvents().list("resourceGroupPS1", "vault1", null, com.azure.core.util.Context.NONE);
    }
}
