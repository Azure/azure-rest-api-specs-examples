
/**
 * Samples for Schemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_Schema.json
     */
    /**
     * Sample code: Schemas_Get.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void schemasGet(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemas().getWithResponse("myResourceGroup", "my-schema-registry", "my-schema",
            com.azure.core.util.Context.NONE);
    }
}
