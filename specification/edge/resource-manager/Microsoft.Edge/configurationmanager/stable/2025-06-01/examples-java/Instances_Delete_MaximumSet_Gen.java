
/**
 * Samples for Instances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Instances_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instancesDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instances().delete("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
