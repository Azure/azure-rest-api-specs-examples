
/**
 * Samples for DynamicSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemas_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemas_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemasGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemas().getWithResponse("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
