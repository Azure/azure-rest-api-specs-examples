
/**
 * Samples for Instances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Instances_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        instancesGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instances().getWithResponse("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
