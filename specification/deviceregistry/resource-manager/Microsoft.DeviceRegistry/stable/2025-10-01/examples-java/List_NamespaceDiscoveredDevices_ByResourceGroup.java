
/**
 * Samples for NamespaceDiscoveredDevices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_NamespaceDiscoveredDevices_ByResourceGroup.json
     */
    /**
     * Sample code: List_NamespaceDiscoveredDevices_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listNamespaceDiscoveredDevicesByResourceGroup(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredDevices().listByResourceGroup("myResourceGroup", "my-namespace-1",
            com.azure.core.util.Context.NONE);
    }
}
