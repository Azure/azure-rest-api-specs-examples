import com.azure.core.util.Context;

/** Samples for DiskPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_ListBySubscription.json
     */
    /**
     * Sample code: List Disk Pools by subscription.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listDiskPoolsBySubscription(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().list(Context.NONE);
    }
}
