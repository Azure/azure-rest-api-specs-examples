
/**
 * Samples for StorageAssets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageAssets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageAssets_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void storageAssetsGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageAssets().getWithResponse("rgdiscovery", "880d16db0ce8a7a846", "c2e4ac21c9ead37737",
            com.azure.core.util.Context.NONE);
    }
}
