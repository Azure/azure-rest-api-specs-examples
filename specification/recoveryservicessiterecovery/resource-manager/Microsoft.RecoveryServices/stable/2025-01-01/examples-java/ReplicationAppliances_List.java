
/**
 * Samples for ReplicationAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationAppliances_List.json
     */
    /**
     * Sample code: Gets the list of appliances.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheListOfAppliances(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationAppliances().list("resourceGroupPS1", "vault1", null, com.azure.core.util.Context.NONE);
    }
}
