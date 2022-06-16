import com.azure.core.util.Context;

/** Samples for UsageModels List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/UsageModels_List.json
     */
    /**
     * Sample code: UsageModels_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void usageModelsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.usageModels().list(Context.NONE);
    }
}
