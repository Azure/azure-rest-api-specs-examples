
/**
 * Samples for SolutionTemplates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplates_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplates_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplatesListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplates().list(com.azure.core.util.Context.NONE);
    }
}
