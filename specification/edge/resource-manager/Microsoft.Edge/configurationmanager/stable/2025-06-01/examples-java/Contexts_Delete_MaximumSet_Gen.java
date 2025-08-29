
/**
 * Samples for Contexts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Contexts_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Contexts_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        contextsDeleteMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.contexts().delete("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
