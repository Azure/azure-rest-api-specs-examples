import com.azure.core.util.Context;

/** Samples for StorageTargets ListByCache. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_ListByCache.json
     */
    /**
     * Sample code: StorageTargets_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().listByCache("scgroup", "sc1", Context.NONE);
    }
}
