import com.azure.core.util.Context;

/** Samples for StorageTargets DnsRefresh. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_DnsRefresh.json
     */
    /**
     * Sample code: Caches_DnsRefresh.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDnsRefresh(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().dnsRefresh("scgroup", "sc", "st1", Context.NONE);
    }
}
