
/**
 * Samples for Schemas ListBySchemaRegistry.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_Schemas_SchemaRegistry.json
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
