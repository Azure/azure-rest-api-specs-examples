
/**
 * Samples for Schemas List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Schemas_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Schemas_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemasListBySubscriptionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemas().list(com.azure.core.util.Context.NONE);
    }
}
