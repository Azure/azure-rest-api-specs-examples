
/**
 * Samples for NamespaceDiscoveredDevices Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_NamespaceDiscoveredDevice.json
     */
    /**
     * Sample code: Get_NamespaceDiscoveredDevice.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        getNamespaceDiscoveredDevice(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredDevices().getWithResponse("myResourceGroup", "my-namespace-1",
            "my-discovereddevice-1", com.azure.core.util.Context.NONE);
    }
}
