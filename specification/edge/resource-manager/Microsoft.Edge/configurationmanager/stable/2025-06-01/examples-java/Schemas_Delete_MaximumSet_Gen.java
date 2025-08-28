
/**
 * Samples for Schemas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Schemas_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Schemas_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        schemasDeleteMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemas().delete("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
