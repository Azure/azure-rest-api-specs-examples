import com.azure.core.util.Context;

/** Samples for DiskPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Delete.json
     */
    /**
     * Sample code: Delete Disk pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void deleteDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().delete("myResourceGroup", "myDiskPool", Context.NONE);
    }
}
