
import com.azure.resourcemanager.workloadorchestration.fluent.models.SchemaVersionInner;
import com.azure.resourcemanager.workloadorchestration.fluent.models.SchemaVersionWithUpdateTypeInner;
import com.azure.resourcemanager.workloadorchestration.models.SchemaVersionProperties;
import com.azure.resourcemanager.workloadorchestration.models.UpdateType;

/**
 * Samples for Schemas CreateVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Schemas_CreateVersion_MaximumSet_Gen.json
     */
    /**
     * Sample code: Schemas_CreateVersion_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void schemasCreateVersionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.schemas().createVersion("rgconfigurationmanager", "testname",
            new SchemaVersionWithUpdateTypeInner().withUpdateType(UpdateType.MAJOR).withVersion("1.0.0")
                .withSchemaVersion(
                    new SchemaVersionInner().withProperties(new SchemaVersionProperties().withValue("uiaqdwsi"))),
            com.azure.core.util.Context.NONE);
    }
}
