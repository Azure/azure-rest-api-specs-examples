
/**
 * Samples for SchemaVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaVersions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaVersions_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaVersionsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaVersions().delete("rgconfigurationmanager", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
