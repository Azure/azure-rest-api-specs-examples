
import com.azure.resourcemanager.resiliencemanagement.models.RecoveryActionRequest;

/**
 * Samples for RecoveryJobs Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryJobs_Cancel_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryJobs_Cancel_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryJobsCancelMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryJobs().cancel("sampleServiceGroup", "qmn", "samplePlanName",
            "c56888ef-9ced-4001-a6d4-7145a0309bdb",
            new RecoveryActionRequest().withDescription("Cancelling the recovery job due to user request"),
            com.azure.core.util.Context.NONE);
    }
}
