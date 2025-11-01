
/**
 * Samples for Schemas ListBySchemaRegistry.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_Schemas_BySchemaRegistry.json
     */
    /**
     * Sample code: List_Schemas_SchemaRegistry.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listSchemasSchemaRegistry(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemas().listBySchemaRegistry("myResourceGroup", "my-schema-registry",
            com.azure.core.util.Context.NONE);
    }
}
