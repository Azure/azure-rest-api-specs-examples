
/**
 * Samples for WorkflowVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/WorkflowVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: WorkflowVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflowVersions().getWithResponse("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
