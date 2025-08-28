
/**
 * Samples for Contexts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Contexts_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Contexts_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void contextsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.contexts().listByResourceGroup("rgconfigurationmanager", com.azure.core.util.Context.NONE);
    }
}
