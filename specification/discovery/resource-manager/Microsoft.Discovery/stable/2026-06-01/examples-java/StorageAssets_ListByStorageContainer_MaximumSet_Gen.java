
/**
 * Samples for StorageAssets ListByStorageContainer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageAssets_ListByStorageContainer_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageAssets_ListByStorageContainer_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        storageAssetsListByStorageContainerMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageAssets().listByStorageContainer("rgdiscovery", "78d6139ad7238f844f",
            com.azure.core.util.Context.NONE);
    }
}
