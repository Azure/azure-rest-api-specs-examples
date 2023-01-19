/** Samples for DiskPools Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Start.json
     */
    /**
     * Sample code: Start Disk Pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void startDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().start("myResourceGroup", "myDiskPool", com.azure.core.util.Context.NONE);
    }
}
