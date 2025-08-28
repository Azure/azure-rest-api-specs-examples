
/**
 * Samples for SchemaVersions ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaVersions_ListBySchema_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaVersions_ListBySchema_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaVersionsListBySchemaMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaVersions().listBySchema("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
