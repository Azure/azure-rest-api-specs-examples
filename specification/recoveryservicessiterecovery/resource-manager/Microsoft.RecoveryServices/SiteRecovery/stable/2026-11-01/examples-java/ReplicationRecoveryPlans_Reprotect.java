
/**
 * Samples for ReplicationRecoveryPlans Reprotect.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationRecoveryPlans_Reprotect.json
     */
    /**
     * Sample code: Execute reprotect of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeReprotectOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().reprotect("resourceGroupPS1", "vault1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
