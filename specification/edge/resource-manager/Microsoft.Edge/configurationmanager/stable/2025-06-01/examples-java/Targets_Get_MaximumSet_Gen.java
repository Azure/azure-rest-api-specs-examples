
/**
 * Samples for Targets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        targetsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().getByResourceGroupWithResponse("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
