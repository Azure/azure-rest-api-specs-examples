
/**
 * Samples for Targets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().list(com.azure.core.util.Context.NONE);
    }
}
