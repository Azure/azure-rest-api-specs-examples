
/**
 * Samples for NamespaceDiscoveredAssets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_NamespaceDiscoveredAsset.json
     */
    /**
     * Sample code: Delete_NamespaceDiscoveredAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        deleteNamespaceDiscoveredAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredAssets().delete("myResourceGroup", "my-namespace-1", "my-discoveredasset-1",
            com.azure.core.util.Context.NONE);
    }
}
