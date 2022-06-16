import com.azure.core.util.Context;

/** Samples for DiskPoolZones List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPoolZones_List.json
     */
    /**
     * Sample code: List Disk Pool Zones.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listDiskPoolZones(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.diskPoolZones().list("eastus", Context.NONE);
    }
}
