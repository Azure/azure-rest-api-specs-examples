/** Samples for DiskPools ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_ListByResourceGroup.json
     */
    /**
     * Sample code: List Disk Pools.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listDiskPools(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
