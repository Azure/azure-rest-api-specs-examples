
import com.azure.resourcemanager.resiliencemanagement.models.TestFailoverCleanupRequest;

/**
 * Samples for RecoveryPlanActions TestFailoverCleanup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_TestFailoverCleanup_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_TestFailoverCleanup_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsTestFailoverCleanupMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().testFailoverCleanup("sampleServiceGroupName", "qmn", "samplePlanName",
            new TestFailoverCleanupRequest().withComments("Test failover clean-up comments"),
            com.azure.core.util.Context.NONE);
    }
}
