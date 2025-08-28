
/**
 * Samples for Targets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        targetsDeleteMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().delete("rgconfigurationmanager", "testname", true, com.azure.core.util.Context.NONE);
    }
}
