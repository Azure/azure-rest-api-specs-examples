
/**
 * Samples for DynamicSchemas ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemas_ListBySchema_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemas_ListBySchema_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemasListBySchemaMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemas().listBySchema("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
