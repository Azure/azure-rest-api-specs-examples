
/**
 * Samples for Contexts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Contexts_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Contexts_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        contextsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.contexts().getByResourceGroupWithResponse("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
