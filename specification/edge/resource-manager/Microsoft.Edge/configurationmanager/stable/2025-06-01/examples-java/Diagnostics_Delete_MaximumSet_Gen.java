
/**
 * Samples for Diagnostics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Diagnostics_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Diagnostics_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void diagnosticsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.diagnostics().delete("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
