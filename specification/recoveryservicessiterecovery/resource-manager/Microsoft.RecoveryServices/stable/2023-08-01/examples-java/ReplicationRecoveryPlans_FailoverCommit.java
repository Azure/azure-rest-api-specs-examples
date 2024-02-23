
/**
 * Samples for ReplicationRecoveryPlans FailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationRecoveryPlans_FailoverCommit.json
     */
    /**
     * Sample code: Execute commit failover of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeCommitFailoverOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().failoverCommit("vault1", "resourceGroupPS1", "RPtest1",
            com.azure.core.util.Context.NONE);
    }
}
