import com.azure.core.util.Context;

/** Samples for DiskPools Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Upgrade.json
     */
    /**
     * Sample code: Upgrade Disk Pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void upgradeDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPools().upgrade("myResourceGroup", "myDiskPool", Context.NONE);
    }
}
