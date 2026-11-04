
/**
 * Samples for ReplicationRecoveryPlans FailoverCancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationRecoveryPlans_FailoverCancel.json
     */
    /**
     * Sample code: Execute cancel failover of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeCancelFailoverOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().failoverCancel("resourceGroupPS1", "vault1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
