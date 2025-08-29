
/**
 * Samples for Workflows Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Workflows_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workflows_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        workflowsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflows().getWithResponse("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
