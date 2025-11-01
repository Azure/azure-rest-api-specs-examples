
/**
 * Samples for SchemaVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_SchemaVersion.json
     */
    /**
     * Sample code: Delete_SchemaVersion.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteSchemaVersion(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaVersions().delete("myResourceGroup", "my-schema-registry", "my-schema", "1",
            com.azure.core.util.Context.NONE);
    }
}
