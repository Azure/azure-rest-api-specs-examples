
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanTestFailoverCleanupInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanTestFailoverCleanupInputProperties;

/**
 * Samples for ReplicationRecoveryPlans TestFailoverCleanup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationRecoveryPlans_TestFailoverCleanup.json
     */
    /**
     * Sample code: Execute test failover cleanup of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeTestFailoverCleanupOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().testFailoverCleanup("vault1", "resourceGroupPS1", "RPtest1",
            new RecoveryPlanTestFailoverCleanupInput().withProperties(
                new RecoveryPlanTestFailoverCleanupInputProperties().withComments("Test Failover Cleanup")),
            com.azure.core.util.Context.NONE);
    }
}
