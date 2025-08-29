
/**
 * Samples for SchemaReferences Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaReferences_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaReferences_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaReferencesGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaReferences().getWithResponse("jdvtghygpz", "testname", com.azure.core.util.Context.NONE);
    }
}
