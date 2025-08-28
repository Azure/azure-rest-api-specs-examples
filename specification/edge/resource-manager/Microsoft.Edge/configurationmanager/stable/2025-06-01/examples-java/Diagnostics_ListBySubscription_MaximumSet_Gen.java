
/**
 * Samples for Diagnostics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Diagnostics_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Diagnostics_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void diagnosticsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.diagnostics().list(com.azure.core.util.Context.NONE);
    }
}
