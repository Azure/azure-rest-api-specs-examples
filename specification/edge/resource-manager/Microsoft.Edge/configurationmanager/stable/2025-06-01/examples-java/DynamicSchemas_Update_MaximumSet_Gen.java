
import com.azure.resourcemanager.workloadorchestration.models.DynamicSchema;
import com.azure.resourcemanager.workloadorchestration.models.DynamicSchemaProperties;

/**
 * Samples for DynamicSchemas Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DynamicSchemas_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: DynamicSchemas_Update_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void dynamicSchemasUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        DynamicSchema resource = manager.dynamicSchemas()
            .getWithResponse("rgconfigurationmanager", "testname", "testname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new DynamicSchemaProperties()).apply();
    }
}
