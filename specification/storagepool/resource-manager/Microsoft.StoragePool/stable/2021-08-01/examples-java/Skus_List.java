
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/Skus_List.json
     */
    /**
     * Sample code: List Disk Pool Skus.
     * 
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listDiskPoolSkus(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.resourceSkus().list("eastus", com.azure.core.util.Context.NONE);
    }
}
