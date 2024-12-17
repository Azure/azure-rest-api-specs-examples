
/**
 * Samples for SchemaRegistries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Delete_SchemaRegistry.json
     */
    /**
     * Sample code: Delete_SchemaRegistry.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteSchemaRegistry(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaRegistries().delete("myResourceGroup", "my-schema-registry", com.azure.core.util.Context.NONE);
    }
}
