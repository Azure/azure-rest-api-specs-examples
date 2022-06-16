import com.azure.core.util.Context;

/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Skus_List.json
     */
    /**
     * Sample code: Skus_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void skusList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.skus().list(Context.NONE);
    }
}
