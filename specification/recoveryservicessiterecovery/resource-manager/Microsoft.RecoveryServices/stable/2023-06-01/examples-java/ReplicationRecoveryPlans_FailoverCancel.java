/** Samples for ReplicationRecoveryPlans FailoverCancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryPlans_FailoverCancel.json
     */
    /**
     * Sample code: Execute cancel failover of the recovery plan.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeCancelFailoverOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationRecoveryPlans()
            .failoverCancel("vault1", "resourceGroupPS1", "RPtest1", com.azure.core.util.Context.NONE);
    }
}
