
import com.azure.resourcemanager.workloadorchestration.models.SchemaVersionProperties;

/**
 * Samples for SchemaVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaVersions_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaVersions_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaVersionsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemaVersions().define("1.0.0").withExistingSchema("rgconfigurationmanager", "testname")
            .withProperties(new SchemaVersionProperties().withValue("uiaqdwsi")).create();
    }
}
