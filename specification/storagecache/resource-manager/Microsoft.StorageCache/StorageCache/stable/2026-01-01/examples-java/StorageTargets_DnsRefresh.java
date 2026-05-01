
/**
 * Samples for StorageTargets DnsRefresh.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/StorageTargets_DnsRefresh.json
     */
    /**
     * Sample code: Caches_DnsRefresh.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDnsRefresh(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().dnsRefresh("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
