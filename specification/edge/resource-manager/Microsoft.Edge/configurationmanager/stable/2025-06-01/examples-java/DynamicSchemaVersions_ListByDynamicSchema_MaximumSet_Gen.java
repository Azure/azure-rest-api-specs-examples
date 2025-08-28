
/**
 * Samples for DynamicSchemaVersions ListByDynamicSchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemaVersions_ListByDynamicSchema_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemaVersions_ListByDynamicSchema_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemaVersionsListByDynamicSchemaMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemaVersions().listByDynamicSchema("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
