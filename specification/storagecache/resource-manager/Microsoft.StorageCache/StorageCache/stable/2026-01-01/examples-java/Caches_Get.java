
/**
 * Samples for Caches GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_Get.json
     */
    /**
     * Sample code: Caches_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().getByResourceGroupWithResponse("scgroup", "sc1", com.azure.core.util.Context.NONE);
    }
}
