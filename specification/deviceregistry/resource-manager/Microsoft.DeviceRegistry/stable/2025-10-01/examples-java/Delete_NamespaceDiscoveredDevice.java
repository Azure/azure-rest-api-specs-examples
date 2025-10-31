
/**
 * Samples for NamespaceDiscoveredDevices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_NamespaceDiscoveredDevice.json
     */
    /**
     * Sample code: Delete_NamespaceDiscoveredDevice.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        deleteNamespaceDiscoveredDevice(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredDevices().delete("myResourceGroup", "my-namespace-1", "my-discovereddevice-1",
            com.azure.core.util.Context.NONE);
    }
}
