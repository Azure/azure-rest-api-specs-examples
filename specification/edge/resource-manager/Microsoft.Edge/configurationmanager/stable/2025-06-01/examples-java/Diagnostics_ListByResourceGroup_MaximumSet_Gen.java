
/**
 * Samples for Diagnostics ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Diagnostics_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Diagnostics_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void diagnosticsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.diagnostics().listByResourceGroup("rgconfigurationmanager", com.azure.core.util.Context.NONE);
    }
}
