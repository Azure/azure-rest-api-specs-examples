
/**
 * Samples for NamespaceDevices Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Get_NamespaceDevice.json
     */
    /**
     * Sample code: Get_NamespaceDevice.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getNamespaceDevice(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().getWithResponse("myResourceGroup", "my-namespace-1", "my-device-name",
            com.azure.core.util.Context.NONE);
    }
}
