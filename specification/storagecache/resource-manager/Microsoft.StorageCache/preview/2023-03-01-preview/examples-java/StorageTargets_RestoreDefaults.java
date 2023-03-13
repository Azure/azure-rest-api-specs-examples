/** Samples for StorageTargets RestoreDefaults. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-03-01-preview/examples/StorageTargets_RestoreDefaults.json
     */
    /**
     * Sample code: StorageTargets_RestoreDefaults.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsRestoreDefaults(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().restoreDefaults("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
