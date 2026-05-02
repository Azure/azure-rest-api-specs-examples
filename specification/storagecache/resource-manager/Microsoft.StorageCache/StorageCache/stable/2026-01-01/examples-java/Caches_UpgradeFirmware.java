
/**
 * Samples for Caches UpgradeFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_UpgradeFirmware.json
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
