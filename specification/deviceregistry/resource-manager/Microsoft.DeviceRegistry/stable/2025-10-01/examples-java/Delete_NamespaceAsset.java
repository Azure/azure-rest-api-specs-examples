
/**
 * Samples for NamespaceAssets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_NamespaceAsset.json
     */
    /**
     * Sample code: Delete_NamespaceAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteNamespaceAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceAssets().delete("myResourceGroup", "adr-namespace-gbk0925-n01", "adr-asset-gbk0925-n01",
            com.azure.core.util.Context.NONE);
    }
}
