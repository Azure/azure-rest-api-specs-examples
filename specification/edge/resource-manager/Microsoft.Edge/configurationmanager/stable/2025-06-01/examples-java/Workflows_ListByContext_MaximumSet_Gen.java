
/**
 * Samples for Workflows ListByContext.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Workflows_ListByContext_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workflows_ListByContext_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowsListByContextMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflows().listByContext("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
