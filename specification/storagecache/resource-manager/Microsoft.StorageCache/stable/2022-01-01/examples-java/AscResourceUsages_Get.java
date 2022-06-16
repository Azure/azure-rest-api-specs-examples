import com.azure.core.util.Context;

/** Samples for AscUsages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/AscResourceUsages_Get.json
     */
    /**
     * Sample code: AscUsages_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void ascUsagesList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.ascUsages().list("eastus", Context.NONE);
    }
}
