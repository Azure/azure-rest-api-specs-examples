
import com.azure.resourcemanager.workloadorchestration.models.SchemaVersion;
import com.azure.resourcemanager.workloadorchestration.models.SchemaVersionProperties;

/**
 * Samples for SchemaVersions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SchemaVersions_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: SchemaVersions_Update_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemaVersionsUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        SchemaVersion resource = manager.schemaVersions()
            .getWithResponse("rgconfigurationmanager", "testname", "1.0.0", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new SchemaVersionProperties().withValue("muezi")).apply();
    }
}
