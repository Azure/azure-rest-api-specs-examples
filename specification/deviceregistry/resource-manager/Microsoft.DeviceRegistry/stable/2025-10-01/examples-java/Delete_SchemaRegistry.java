
/**
 * Samples for SchemaRegistries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_SchemaRegistry.json
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
