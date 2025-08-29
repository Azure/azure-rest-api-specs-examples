
/**
 * Samples for Contexts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Contexts_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Contexts_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void contextsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.contexts().list(com.azure.core.util.Context.NONE);
    }
}
