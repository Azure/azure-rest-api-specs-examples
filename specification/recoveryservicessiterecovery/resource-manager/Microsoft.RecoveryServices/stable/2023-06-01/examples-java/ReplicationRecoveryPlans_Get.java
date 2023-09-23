/** Samples for ReplicationRecoveryPlans Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryPlans_Get.json
     */
    /**
     * Sample code: Gets the requested recovery plan.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheRequestedRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationRecoveryPlans()
            .getWithResponse("vault1", "resourceGroupPS1", "RPtest1", com.azure.core.util.Context.NONE);
    }
}
