
/**
 * Samples for ReplicationProtectionIntents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionIntents_Get.json
     */
    /**
     * Sample code: Gets the details of a Replication protection intent item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAReplicationProtectionIntentItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionIntents().getWithResponse("resourceGroupPS1", "vault1", "vm1",
            com.azure.core.util.Context.NONE);
    }
}
