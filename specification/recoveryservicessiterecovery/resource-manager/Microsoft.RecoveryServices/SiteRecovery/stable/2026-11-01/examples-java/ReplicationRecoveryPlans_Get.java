
/**
 * Samples for ReplicationRecoveryPlans Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationRecoveryPlans_Get.json
     */
    /**
     * Sample code: Gets the requested recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheRequestedRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().getWithResponse("resourceGroupPS1", "vault1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
