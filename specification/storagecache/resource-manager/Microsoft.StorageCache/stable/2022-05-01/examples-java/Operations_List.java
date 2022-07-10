import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void operationsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.operations().list(Context.NONE);
    }
}
