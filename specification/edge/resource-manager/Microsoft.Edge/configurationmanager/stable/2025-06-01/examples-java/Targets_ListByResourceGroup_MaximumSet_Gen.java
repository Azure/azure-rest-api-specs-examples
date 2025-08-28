
/**
 * Samples for Targets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().listByResourceGroup("rgconfigurationmanager", com.azure.core.util.Context.NONE);
    }
}
