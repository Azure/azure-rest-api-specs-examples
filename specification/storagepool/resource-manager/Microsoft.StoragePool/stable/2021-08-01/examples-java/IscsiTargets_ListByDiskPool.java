/** Samples for IscsiTargets ListByDiskPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_ListByDiskPool.json
     */
    /**
     * Sample code: List Disk Pools by Resource Group.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listDiskPoolsByResourceGroup(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.iscsiTargets().listByDiskPool("myResourceGroup", "myDiskPool", com.azure.core.util.Context.NONE);
    }
}
