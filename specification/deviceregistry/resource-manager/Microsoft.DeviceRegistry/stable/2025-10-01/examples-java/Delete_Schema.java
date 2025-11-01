
/**
 * Samples for Schemas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_Schema.json
     */
    /**
     * Sample code: Delete_Schema.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteSchema(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemas().delete("myResourceGroup", "my-schema-registry", "my-schema",
            com.azure.core.util.Context.NONE);
    }
}
