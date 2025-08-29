
/**
 * Samples for SchemaVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaVersions().getWithResponse("rgconfigurationmanager", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
