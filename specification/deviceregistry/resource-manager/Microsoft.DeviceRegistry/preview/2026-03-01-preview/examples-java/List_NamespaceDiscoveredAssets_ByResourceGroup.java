
/**
 * Samples for NamespaceDiscoveredAssets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/List_NamespaceDiscoveredAssets_ByResourceGroup.json
     */
    /**
     * Sample code: List_NamespaceDiscoveredAssets_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listNamespaceDiscoveredAssetsByResourceGroup(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDiscoveredAssets().listByResourceGroup("myResourceGroup", "my-namespace-1",
            com.azure.core.util.Context.NONE);
    }
}
