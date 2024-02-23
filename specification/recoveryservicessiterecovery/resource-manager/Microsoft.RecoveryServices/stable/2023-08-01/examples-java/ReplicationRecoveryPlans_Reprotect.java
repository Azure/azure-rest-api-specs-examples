
/**
 * Samples for ReplicationRecoveryPlans Reprotect.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationRecoveryPlans_Reprotect.json
     */
    /**
     * Sample code: Execute reprotect of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeReprotectOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().reprotect("vault1", "resourceGroupPS1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
