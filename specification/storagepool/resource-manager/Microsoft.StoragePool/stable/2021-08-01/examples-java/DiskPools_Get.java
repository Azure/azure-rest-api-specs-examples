/** Samples for DiskPools GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Get.json
     */
    /**
     * Sample code: Get Disk pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void getDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager
            .diskPools()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskPool", com.azure.core.util.Context.NONE);
    }
}
