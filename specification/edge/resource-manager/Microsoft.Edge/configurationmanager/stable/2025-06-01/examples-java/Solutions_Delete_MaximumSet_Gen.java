
/**
 * Samples for Solutions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Solutions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Solutions_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutions().delete("rgconfigurationmanager", "testname", "testname", com.azure.core.util.Context.NONE);
    }
}
