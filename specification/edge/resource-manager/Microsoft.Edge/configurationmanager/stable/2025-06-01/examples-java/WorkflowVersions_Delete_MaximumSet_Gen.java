
/**
 * Samples for WorkflowVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/WorkflowVersions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkflowVersions_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowVersionsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflowVersions().delete("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
