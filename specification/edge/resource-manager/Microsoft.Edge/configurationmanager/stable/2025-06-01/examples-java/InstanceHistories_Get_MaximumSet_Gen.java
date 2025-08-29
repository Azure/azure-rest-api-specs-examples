
/**
 * Samples for InstanceHistories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/InstanceHistories_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: InstanceHistories_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instanceHistoriesGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instanceHistories().getWithResponse("rgconfigurationmanager", "testname", "testname", "testname",
            "testname", com.azure.core.util.Context.NONE);
    }
}
