import com.azure.core.util.Context;

/** Samples for DiskPools Deallocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Deallocate.json
     */
    /**
     * Sample code: Deallocate Disk Pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void deallocateDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().deallocate("myResourceGroup", "myDiskPool", Context.NONE);
    }
}
