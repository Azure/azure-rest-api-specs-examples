
/**
 * Samples for Solutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Solutions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Solutions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        solutionsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutions().getWithResponse("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
