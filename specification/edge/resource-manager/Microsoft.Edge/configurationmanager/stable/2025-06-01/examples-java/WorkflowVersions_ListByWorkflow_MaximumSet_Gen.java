
/**
 * Samples for WorkflowVersions ListByWorkflow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/WorkflowVersions_ListByWorkflow_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkflowVersions_ListByWorkflow_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowVersionsListByWorkflowMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflowVersions().listByWorkflow("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
