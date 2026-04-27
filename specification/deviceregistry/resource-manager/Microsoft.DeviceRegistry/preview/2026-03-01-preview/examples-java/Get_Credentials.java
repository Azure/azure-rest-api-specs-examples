
/**
 * Samples for Credentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Get_Credentials.json
     */
    /**
     * Sample code: Get_Credentials.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getCredentials(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.credentials().getWithResponse("rgdeviceregistry", "mynamespace", com.azure.core.util.Context.NONE);
    }
}
