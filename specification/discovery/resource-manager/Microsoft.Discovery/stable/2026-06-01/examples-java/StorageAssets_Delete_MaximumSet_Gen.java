
/**
 * Samples for StorageAssets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageAssets_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageAssets_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void storageAssetsDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageAssets().delete("rgdiscovery", "86f9743316a64d577a", "7255aed4c052c5a165",
            com.azure.core.util.Context.NONE);
    }
}
