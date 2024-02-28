
/**
 * Samples for Caches UpgradeFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * Caches_UpgradeFirmware.json
     */
    /**
     * Sample code: Caches_UpgradeFirmware.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesUpgradeFirmware(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().upgradeFirmware("scgroup", "sc1", com.azure.core.util.Context.NONE);
    }
}
