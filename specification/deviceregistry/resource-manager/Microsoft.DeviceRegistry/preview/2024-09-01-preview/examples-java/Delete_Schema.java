
/**
 * Samples for Schemas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Delete_Schema.json
     */
    /**
     * Sample code: Delete_Schema.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteSchema(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemas().deleteWithResponse("myResourceGroup", "my-schema-registry", "my-schema",
            com.azure.core.util.Context.NONE);
    }
}
