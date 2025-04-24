
/**
 * Samples for ReplicationRecoveryPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationRecoveryPlans_Delete.json
     */
    /**
     * Sample code: Deletes the specified recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void deletesTheSpecifiedRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().delete("resourceGroupPS1", "vault1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
