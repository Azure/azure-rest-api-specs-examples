
/**
 * Samples for NamespaceDevices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_NamespaceDevice.json
     */
    /**
     * Sample code: Delete_NamespaceDevice.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteNamespaceDevice(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().delete("myResourceGroup", "adr-namespace-gbk0925-n01", "adr-device-gbk0925-n01",
            com.azure.core.util.Context.NONE);
    }
}
