
/**
 * Samples for NamespaceDiscoveredAssets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_NamespaceDiscoveredAsset.json
     */
    /**
     * Sample code: Get_NamespaceDiscoveredAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        getNamespaceDiscoveredAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredAssets().getWithResponse("myResourceGroup", "my-namespace-1", "my-discoveredasset-1",
            com.azure.core.util.Context.NONE);
    }
}
