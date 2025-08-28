
/**
 * Samples for DynamicSchemaVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemaVersions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemaVersions_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemaVersionsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemaVersions().delete("rgconfigurationmanager", "testname", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
