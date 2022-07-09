import com.azure.core.util.Context;

/** Samples for Caches GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/Caches_Get.json
     */
    /**
     * Sample code: Caches_Get.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().getByResourceGroupWithResponse("scgroup", "sc1", Context.NONE);
    }
}
