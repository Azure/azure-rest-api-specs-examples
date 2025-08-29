
import com.azure.resourcemanager.workloadorchestration.models.SchemaVersionProperties;

/**
 * Samples for DynamicSchemaVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemaVersions_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemaVersions_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemaVersionsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.dynamicSchemaVersions().define("1.0.0")
            .withExistingDynamicSchema("rgconfigurationmanager", "testname", "testname")
            .withProperties(new SchemaVersionProperties().withValue("uiaqdwsi")).create();
    }
}
