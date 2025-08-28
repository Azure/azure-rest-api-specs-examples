
/**
 * Samples for DynamicSchemas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemas_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemas_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemasDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemas().delete("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
