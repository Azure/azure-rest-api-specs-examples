
/**
 * Samples for SchemaVersions ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/List_SchemaVersions_Schema.json
     */
    /**
     * Sample code: List_SchemaVersions_Schema.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listSchemaVersionsSchema(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaVersions().listBySchema("myResourceGroup", "my-schema-registry", "my-schema",
            com.azure.core.util.Context.NONE);
    }
}
