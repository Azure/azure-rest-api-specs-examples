
import com.azure.resourcemanager.workloadorchestration.models.DynamicSchemaProperties;

/**
 * Samples for DynamicSchemas CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemas_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemas_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemasCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemas().define("testname").withExistingSchema("rgconfigurationmanager", "testname")
            .withProperties(new DynamicSchemaProperties()).create();
    }
}
