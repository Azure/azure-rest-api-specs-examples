
/**
 * Samples for NamespaceAssets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_NamespaceAssets_ByResourceGroup.json
     */
    /**
     * Sample code: List_NamespaceAssets_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listNamespaceAssetsByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceAssets().listByResourceGroup("myResourceGroup", "adr-namespace-gbk0925-n01",
            com.azure.core.util.Context.NONE);
    }
}
