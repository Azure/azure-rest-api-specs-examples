
/**
 * Samples for NamespaceAssets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_NamespaceAsset.json
     */
    /**
     * Sample code: Get_NamespaceAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getNamespaceAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceAssets().getWithResponse("myResourceGroup", "my-namespace-1", "my-asset-1",
            com.azure.core.util.Context.NONE);
    }
}
