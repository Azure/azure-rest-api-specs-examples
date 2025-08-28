
/**
 * Samples for ConfigTemplates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplates_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplates_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplatesListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplates().list(com.azure.core.util.Context.NONE);
    }
}
