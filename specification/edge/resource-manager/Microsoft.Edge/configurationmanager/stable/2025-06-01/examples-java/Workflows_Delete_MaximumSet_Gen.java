
/**
 * Samples for Workflows Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Workflows_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workflows_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflows().delete("rgconfigurationmanager", "testname", "testname", com.azure.core.util.Context.NONE);
    }
}
