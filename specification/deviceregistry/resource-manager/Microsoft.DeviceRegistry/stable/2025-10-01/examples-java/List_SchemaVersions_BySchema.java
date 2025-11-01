
/**
 * Samples for SchemaVersions ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_SchemaVersions_BySchema.json
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
