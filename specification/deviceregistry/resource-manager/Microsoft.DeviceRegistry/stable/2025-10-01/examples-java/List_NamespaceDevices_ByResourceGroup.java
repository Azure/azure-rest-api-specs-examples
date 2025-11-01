
/**
 * Samples for NamespaceDevices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_NamespaceDevices_ByResourceGroup.json
     */
    /**
     * Sample code: List_NamespaceDevices_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listNamespaceDevicesByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().listByResourceGroup("myResourceGroup", "adr-namespace-gbk0925-n01",
            com.azure.core.util.Context.NONE);
    }
}
