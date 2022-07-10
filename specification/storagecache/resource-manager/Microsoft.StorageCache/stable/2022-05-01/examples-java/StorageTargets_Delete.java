import com.azure.core.util.Context;

/** Samples for StorageTargets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/StorageTargets_Delete.json
     */
    /**
     * Sample code: StorageTargets_Delete.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().delete("scgroup", "sc1", "st1", null, Context.NONE);
    }
}
