
/**
 * Samples for Diagnostics GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Diagnostics_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Diagnostics_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        diagnosticsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.diagnostics().getByResourceGroupWithResponse("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
