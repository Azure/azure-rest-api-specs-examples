import com.azure.core.util.Context;

/** Samples for Caches ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/Caches_ListByResourceGroup.json
     */
    /**
     * Sample code: Caches_ListByResourceGroup.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesListByResourceGroup(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().listByResourceGroup("scgroup", Context.NONE);
    }
}
