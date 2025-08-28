
/**
 * Samples for InstanceHistories ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/InstanceHistories_ListByInstance_MaximumSet_Gen.json
     */
    /**
     * Sample code: InstanceHistories_ListByInstance_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instanceHistoriesListByInstanceMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instanceHistories().listByInstance("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
