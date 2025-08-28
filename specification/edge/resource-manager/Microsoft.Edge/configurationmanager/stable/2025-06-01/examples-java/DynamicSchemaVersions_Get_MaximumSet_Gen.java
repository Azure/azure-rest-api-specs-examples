
/**
 * Samples for DynamicSchemaVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemaVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemaVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemaVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemaVersions().getWithResponse("rgconfigurationmanager", "testname", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
