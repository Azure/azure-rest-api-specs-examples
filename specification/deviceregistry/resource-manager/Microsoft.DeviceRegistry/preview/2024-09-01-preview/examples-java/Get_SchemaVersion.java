
/**
 * Samples for SchemaVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_SchemaVersion.json
     */
    /**
     * Sample code: Get_SchemaVersion.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getSchemaVersion(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaVersions().getWithResponse("myResourceGroup", "my-schema-registry", "my-schema", "1",
            com.azure.core.util.Context.NONE);
    }
}
